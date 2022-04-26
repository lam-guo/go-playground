package playground

import (
	"fmt"
	"testing"
)

// defer exec at the end of the function
// defer: first in last out
func TestDeferInLoop(t *testing.T) {
	for i := 0; i < 10; i++ {
		defer func(num int) {
			fmt.Println("defer:", num)
		}(i)
		fmt.Println(i)
	}
}
