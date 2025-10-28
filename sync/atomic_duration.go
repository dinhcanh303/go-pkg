package sync

import (
	"sync/atomic"
	"time"
)

type AtomicDuration int64

func NewAtomicDuration() *AtomicDuration {
	return new(AtomicDuration)
}

func ForAtomicDuration(v time.Duration) *AtomicDuration {
	d := NewAtomicDuration()
	d.Set(v)
	return d
}

func (ad *AtomicDuration) Set(v time.Duration) {
	atomic.StoreInt64((*int64)(ad), int64(v))
}

func (ad *AtomicDuration) Load() time.Duration {
	return time.Duration(atomic.LoadInt64((*int64)(ad)))
}

func (ad *AtomicDuration) CompareAndSwap(old, new time.Duration) bool {
	return atomic.CompareAndSwapInt64((*int64)(ad), int64(old), int64(new))
}
