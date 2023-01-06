package g_pool

type worker struct {
	task chan f
	// 回收worker的goroutine
	clean chan struct{}
	pool  *pool
}

type f func() error

func (w *worker) run() {
	go func() {
		for {
			select {
			case t := <-w.task:
				t()
				w.pool.PutBackWorker(w)
			case <-w.clean:
				return
			}
		}
	}()
}

func (w *worker) SetTask(t f) {
	w.task <- t
}
