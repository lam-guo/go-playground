package g_pool

type worker struct {
	task chan f
}

type f func() error

func (w *worker) run() {
	go func() {
		for task := range w.task {
			task()
		}
	}()
}
