package sync

import (
	"math"
	"sync/atomic"
)

type AtomicFloat64 int64

func NewAtomicFloat64() *AtomicFloat64 {
	return new(AtomicFloat64)
}

func ForAtomicFloat64(v float64) *AtomicFloat64 {
	f := NewAtomicFloat64()
	f.Set(v)
	return f
}

func (af *AtomicFloat64) Set(v float64) {
	atomic.StoreInt64((*int64)(af), int64(math.Float64bits(v)))
}

func (af *AtomicFloat64) Load() float64 {
	return math.Float64frombits(uint64(atomic.LoadInt64((*int64)(af))))
}

func (af *AtomicFloat64) CompareAndSwap(old, new float64) bool {
	return atomic.CompareAndSwapInt64((*int64)(af), int64(math.Float64bits(old)), int64(math.Float64bits(new)))
}

func (f *AtomicFloat64) Add(v float64) float64 {
	for {
		old := f.Load()
		nv := old + v
		if f.CompareAndSwap(old, nv) {
			return nv
		}
	}
}
