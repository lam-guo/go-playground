package loop_test

import (
	"fmt"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestWhileLoop(t *testing.T) {
	n := 0
	/* while(n<5) */
	for n < 5 {
		t.Log(n)
		n++
	}
}

func TestGoLoop(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	eg := new(errgroup.Group)
	for _, v := range arr {
		eg.Go(func() error {
			fmt.Println(v) // 可能打印 7 7 7 7 7 7 7,因为都是对v的引用
			return nil
		})
	}
	eg.Wait()
}

// 闭包写法1
func TestGoLoop1(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	eg := new(errgroup.Group)
	for _, v := range arr {
		func(val int) {
			eg.Go(func() error {
				fmt.Println(val)
				return nil
			})
		}(v)
	}
	eg.Wait()
}

// 闭包写法2
func TestGoLoop2(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	eg := new(errgroup.Group)
	for _, v := range arr {
		val := v
		eg.Go(func() error {
			fmt.Println(val)
			return nil
		})
	}
	eg.Wait()
}
