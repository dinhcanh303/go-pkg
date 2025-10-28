package thread

import "sync"

type RoutineGroup struct {
	waitGroup sync.WaitGroup
}

func NewRoutineGroup() *RoutineGroup {
	return new(RoutineGroup)
}

func (rg *RoutineGroup) Run(fn func()) {
	rg.waitGroup.Add(1)
	go func() {
		defer rg.waitGroup.Done()
		fn()
	}()
}

func (rg *RoutineGroup) RunSafe(fn func()) {
	rg.waitGroup.Add(1)

}

func (rg *RoutineGroup) Wait() {
	rg.waitGroup.Wait()
}
