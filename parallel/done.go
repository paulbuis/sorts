package parallel

import (
	"sync"
)

type Done struct { // implements sync.Locker
	m *sync.Mutex  // must be a *sync.Mutex, not a sync.Mutex,
}                  // since Done may be used as a func return value and hence copied

func NewDone() Done {
	d := Done{m: &sync.Mutex{} }
	d.m.Lock()
	return d
}

func (d Done) Lock() {
	d.m.Lock()
}

func (d Done) Unlock() {
	d.m.Unlock()
}