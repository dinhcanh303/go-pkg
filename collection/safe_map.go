package collection

import "sync"

const (
	maxDeletion = 10000
)

// SafeMap is a thread-safe map with dual dirty maps to optimize read and write operations.
type SafeMap struct {
	lock        sync.RWMutex
	dirtyOld    map[any]any
	dirtyNew    map[any]any
	deletionOld int
	deletionNew int
}

// NewSafeMap creates and returns a new instance of SafeMap.
func NewSafeMap() *SafeMap {
	return &SafeMap{
		dirtyOld: make(map[any]any),
		dirtyNew: make(map[any]any),
	}
}

// Get retrieves the value associated with the given key.
func (sm *SafeMap) Get(key any) (any, bool) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	if val, ok := sm.dirtyOld[key]; ok {
		return val, true
	}
	val, ok := sm.dirtyNew[key]
	return val, ok
}

// Set sets the value for the given key in the SafeMap.
func (sm *SafeMap) Set(key, value any) {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	if sm.deletionOld <= maxDeletion {
		if _, ok := sm.dirtyNew[key]; ok {
			delete(sm.dirtyNew, key)
			sm.deletionNew++
		}
		sm.dirtyOld[key] = value
	} else {
		if _, ok := sm.dirtyOld[key]; ok {
			delete(sm.dirtyOld, key)
			sm.deletionOld++
		}
		sm.dirtyNew[key] = value
	}
}

// Size returns the total number of key-value pairs in the SafeMap.
func (sm *SafeMap) Size() int {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	return len(sm.dirtyOld) + len(sm.dirtyNew)
}

// Range iterates over all key-value pairs in the SafeMap and applies the given function.
func (sm *SafeMap) Range(f func(key, value any) bool) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()

	for k, v := range sm.dirtyOld {
		if !f(k, v) {
			return
		}
	}
	for k, v := range sm.dirtyNew {
		if !f(k, v) {
			return
		}
	}
}
