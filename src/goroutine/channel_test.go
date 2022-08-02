package goroutine

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
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
	var wg sync.WaitGroup
	c := make(chan int, len(sources))
	wg.Add(len(sources))
	for _, ch := range sources {
		go func(in <-chan int) {
			c <- <-in
			wg.Done()
		}(ch)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
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

// Done returns a channel that's closed when work done on behalf of this context should be canceled
// 1.Done可能返回nil如果ctx不能被canceld。Done may return nil if this context can never be canceled
// 2.Done调用会返回同样的值。Successive calls to Done return the same value
func TestCtxDoneChan(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(100 * time.Millisecond)
		// 只有cancel执行了，将ctx的done这个channel 关闭了，<-ctx.Done()才不会阻塞，不然<-ctx.Done()会一直阻塞
		fmt.Println("我还没cancel")
		cancel()
		fmt.Println("cancel了")
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

// 关闭了的channel可以读取还在channel里面的值；如果没有值了，取零值
func TestReusableChan(t *testing.T) {
	var closedchan = make(chan int, 5)
	closedchan <- 1
	closedchan <- 1
	closedchan <- 1

	close(closedchan)

	for i := 0; i < 5; i++ {
		fmt.Println(<-closedchan) // 输出 1 1 1 0 0
	}
}

func TestReusableChanForRange(t *testing.T) {
	var closedchan = make(chan int, 5)
	closedchan <- 1
	closedchan <- 1
	closedchan <- 1

	close(closedchan)

	go func() {
		for v := range closedchan {
			fmt.Println(v) // 输出 1 1 1 0 0
		}
	}()

	// for range 一个close的信道会没有输出
	ch := make(chan struct{})
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println(<-ch)
	fmt.Println("finish")
}

// nil channel会一直堵塞
func TestNilChan(t *testing.T) {
	var ch chan int
	ch = nil
	ch <- 1
	fmt.Println(<-ch)
}

func TestClose(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumSenders = 4003

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	// ...
	dataCh := make(chan int, 10)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the receiver of channel
	// dataCh, and its receivers are the
	// senders of channel dataCh.

	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				// Even if stopCh is closed, the first
				// branch in the second select may be
				// still not selected for some loops if
				// the send to dataCh is also unblocked.
				// But this is acceptable for this
				// example, so the first select block
				// above can be omitted.
				select {
				case <-stopCh:
					return
				case dataCh <- rand.Intn(Max):
					log.Println("not")
				}
			}
		}()
	}

	// the receiver
	go func() {
		defer wgReceivers.Done()

		for value := range dataCh {
			if value == Max-1 {
				log.Print("close")
				// The receiver of channel dataCh is
				// also the sender of stopCh. It is
				// safe to close the stop channel here.
				close(stopCh)
				return
			}

			log.Println(value)
		}
	}()

	// ...
	wgReceivers.Wait()
}
