package thread

import "sync"

// RoutineGroup is a wrapper around sync.WaitGroup to manage multiple goroutines.
type RoutineGroup struct {
	waitGroup sync.WaitGroup
}

// NewRoutineGroup creates and returns a new instance of RoutineGroup.
func NewRoutineGroup() *RoutineGroup {
	return new(RoutineGroup)
}

// Run starts a new goroutine to execute the given fn and tracks it in the RoutineGroup.
func (rg *RoutineGroup) Run(fn func()) {
	rg.waitGroup.Add(1)
	go func() {
		defer rg.waitGroup.Done()
		fn()
	}()
}

// RunSafe starts a new goroutine to execute the given fn safely (recovering from panics) and tracks it in the RoutineGroup.
func (rg *RoutineGroup) RunSafe(fn func()) {
	rg.waitGroup.Add(1)

}

// Wait blocks until all goroutines in the RoutineGroup have completed.
func (rg *RoutineGroup) Wait() {
	rg.waitGroup.Wait()
}
