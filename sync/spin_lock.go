package sync

import (
	"runtime"
	"sync/atomic"
)

type SpinLock struct {
	lock uint32
}

func (s *SpinLock) Lock() {
	if !s.TryLock() {
		runtime.Gosched()
	}
}

func (s *SpinLock) TryLock() bool {
	return atomic.CompareAndSwapUint32(&s.lock, 0, 1)
}

func (s *SpinLock) Unlock() {
	atomic.StoreUint32(&s.lock, 0)
}
