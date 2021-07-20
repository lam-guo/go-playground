package waitgroup

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func WaitTimeOut(wg *sync.WaitGroup, d time.Duration) bool {
	c := make(chan struct{}, 0)
	go func() {
		wg.Wait()
		c <- struct{}{}
	}()
	tm := time.After(d)
	select {
	case <-tm:
		return true
	case <-c:
		return false
	}
}

func TestWg(t *testing.T) {
	wg := &sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			<-close
			fmt.Println(num)
		}(i, c)
	}
	if WaitTimeOut(wg, 5*time.Second) {
		close(c)
		fmt.Println("timeout")
	}
	time.Sleep(time.Second * 10)
	wg.Done()
}
