package waitgroup

import (
	"fmt"
	"sync"
	"sync/atomic"
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

func TestWg2(t *testing.T) {
	var wg sync.WaitGroup
	var counter int32
	n := 2000
	for i := 0; i < n; i++ {
		go func(i int) {
			wg.Add(1)
			atomic.AddInt32(&counter, 1) // 存在wait()提前触发问题
			// counter += 1 // Dangerous!!
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(counter)
	for i := 0; i < 2; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(counter)
	}
}

func TestWg3(t *testing.T) {
	wg := sync.WaitGroup{}
	var counter int
	var other int32
	hashMap := sync.Map{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		other += 1
		// go func(wg sync.WaitGroup, num int) { // 注意这里传入wg是值传递，相当于新的wg,新的wg一直在Done，但是旧的wg一直在Add
		go func(wg *sync.WaitGroup, num int) {
			hashMap.Store(counter, true)
			counter += 1
			wg.Done()
		}(&wg, i)
	}
	// hashMap.Range(func(key, value interface{}) bool {
	// 	fmt.Println(key)
	// 	return true
	// })
	for i := 0; i < 1000; i++ {
		if v, ok := hashMap.Load(i); !ok {
			fmt.Println(i)
		} else {
			fmt.Println(v, "ok")
		}
	}
	wg.Wait()
	fmt.Println(counter, "----")
	fmt.Println(other, "-----")
}
