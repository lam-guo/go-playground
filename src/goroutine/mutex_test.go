package goroutine

import (
	"sync"
	"testing"
	"time"
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

type M struct {
	sync.RWMutex
	v map[int]struct{}
}

func (m *M) Get() {
	m.RLock()
	defer m.RUnlock()
}

func (m *M) Set(i int) {
	m.Lock()
	defer m.Unlock()
	m.v[i] = struct{}{}
}

func TestDeadLock2(t *testing.T) {
	m := M{
		sync.RWMutex{},
		map[int]struct{}{},
	}
	wg := sync.WaitGroup{}
	j := 10000
	wg.Add(j)
	go func() {
		for j > 0 {
			go func(i int) {
				go m.Get()
				go m.Set(i)
				go m.Get()
				time.Sleep(5 * time.Millisecond)
				wg.Done()
			}(j)
			j--
		}
	}()
	wg.Wait()
}
