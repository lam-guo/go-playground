package advanced

import (
	"sync"
	"sync/atomic"
	"testing"
)

// Lock try lock
type Lock struct {
	lock atomic.Int32
}

// NewLock generate a try lock
func NewLock() Lock {
	var l Lock
	l.lock = atomic.Int32{}
	l.lock.Store(0)
	return l
}

// Lock try lock, return lock result
func (l *Lock) Lock() bool {
	return l.lock.CompareAndSwap(0, 1)
}

// Unlock , Unlock the try lock
func (l *Lock) Unlock() {
	l.lock.Store(0)
}

var counter int

func TestLock(t *testing.T) {
	var l = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock() {
				// log error
				println("lock failed")
				return
			}
			counter++
			l.Unlock()
		}()
	}
	wg.Wait()
}
