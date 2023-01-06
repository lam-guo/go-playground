package g_pool

import (
	"sync"
)

type pool struct {
	// slice存放可用worker
	workers []*worker
	// 池子大小
	size int
	// 正在使用的worker数量
	using int
	lock  sync.Mutex
}

func NewPool(size int) *pool {
	return &pool{
		size:    size,
		lock:    sync.Mutex{},
		workers: nil,
	}
}

func (p *pool) GetWorker() *worker {
	p.lock.Lock()
	usableNum := len(p.workers) - 1
	var needWaiting bool
	if usableNum <= 0 {
		needWaiting = p.using >= p.size
		p.lock.Unlock()
	} else {
		workers := p.workers
		w := workers[usableNum]
		p.workers = workers[:usableNum]
		p.lock.Unlock()
		return w
	}

	if needWaiting {
		for {
			p.lock.Lock()
			usableNum = len(p.workers) - 1
			if usableNum <= 0 {
				p.lock.Unlock()
				continue
			}
			workers := p.workers
			w := workers[usableNum]
			p.workers = workers[:usableNum]
			p.lock.Unlock()
			return w
		}
	}
	w := &worker{
		pool:  p,
		task:  make(chan f, 1),
		clean: make(chan struct{}),
	}
	w.run()
	p.workers = append(p.workers, w)
	return w
}

func (p *pool) PutBackWorker(w *worker) {
	p.lock.Lock()
	p.workers = append(p.workers, w)
	p.lock.Unlock()
}
