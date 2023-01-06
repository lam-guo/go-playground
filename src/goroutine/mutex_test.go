package goroutine

import (
	"sync"
	"testing"
)

func TestDeadLock(t *testing.T) {
	rw := sync.RWMutex{}
	rw.RLock()
	rw.Lock()
	rw.RLock()
	go func() {
		rw.Unlock()
	}()
	go func() {
		rw.RUnlock()
	}()
	go func() {
		rw.RUnlock()
	}()
}
