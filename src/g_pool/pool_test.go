package g_pool

import (
	"fmt"
	"testing"
)

func TestPool(t *testing.T) {
	p := NewPool(100)
	for i := 0; i < 1000; i++ {
		w := p.GetWorker()
		w.SetTask(func() error {
			fmt.Println("任务")
			return nil
		})
	}
}
