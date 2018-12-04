package parallel

import (
	"sync"
)

type TaskFunc func()

type TaskGroup struct {
	wg *sync.WaitGroup
}

func MakeTaskGroup() TaskGroup {
	return TaskGroup{wg: &sync.WaitGroup{}}
}

func (tg TaskGroup)Add(task TaskFunc) {
	tg.wg.Add(1)
	go func() {
		task()
		tg.wg.Done()
	}()
}

func (tg TaskGroup) Wait() {
	tg.wg.Wait()
}