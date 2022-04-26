package goroutine

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {

}

func send(ch chan<- string, message string) {
	fmt.Printf("Sending: %#v\n", message)
	ch <- message
}

func read(ch <-chan string) {
	fmt.Printf("Receiving: %#v\n", <-ch)
}

func main() {
	ch := make(chan string, 1)
	send(ch, "Hello World!")
	read(ch)
}

func TestSendAndReceive(t *testing.T) {
	times := 5
	ch := make(chan int)
	done := make(chan bool)
	go func() {
		for i := 0; i < times; i++ {
			ch <- rand.Intn(5)
		}
	}()
	go func() {
		for i := 0; i < times; i++ {
			fmt.Println(<-ch)
		}
		done <- true
	}()
	<-done
}

func fanin(sources ...<-chan int) <-chan int {
	c := make(chan int, len(sources))
	for _, ch := range sources {
		go func(in <-chan int) {
			c <- <-in
		}(ch)
	}
	return c
}

func TestFanIn(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	ch1 <- 1
	ch2 <- 2
	ch3 <- 3
	c := fanin(ch1, ch2, ch3)
	for i := range c {
		t.Log(i)
	}
}

//  Done returns a channel that's closed when work done on behalf of this context should be canceled
// 1.Done可能返回nil如果ctx不能被canceld。Done may return nil if this context can never be canceled
// 2.Done调用会返回同样的值。Successive calls to Done return the same value
func TestCtxDoneChan(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	go func() {
		// 为了测试ctx.Done()是否能调用多次
		for {
			select {
			case <-ctx.Done():
				fmt.Println("我来1了", ctx.Err())
			}
		}
	}()

	go func() {
		<-ctx.Done()
		fmt.Println("我来2了", ctx.Err())
	}()
	time.Sleep(1 * time.Second)
}