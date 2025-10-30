package string

import (
	"math/rand"
	"sync"
)

// lockedSource is a thread-safe wrapper around rand.Source.
type lockedSource struct {
	lock   sync.Mutex
	source rand.Source
}

// newLockedSource creates a new instance of lockedSource with the given seed.
func newLockedSource(seed int64) *lockedSource {
	return &lockedSource{
		source: rand.NewSource(seed),
	}
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (ls *lockedSource) Int63() int64 {
	ls.lock.Lock()
	defer ls.lock.Unlock()
	return ls.source.Int63()
}

// Seed initializes the generator to a deterministic state.
func (ls *lockedSource) Seed(seed int64) {
	ls.lock.Lock()
	defer ls.lock.Unlock()
	ls.source.Seed(seed)
}
