package parallel

import (
	"sync"
)

// Lock return value to sync and await side-effects of f
func Spawn(f func()) sync.Locker {
	done := NewDone() // created in locked state
	go func() {
		f()
		done.Unlock()
	} ()
	return done
}