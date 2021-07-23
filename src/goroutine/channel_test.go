package goroutine

import (
	"fmt"
	"math/rand"
	"testing"
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
