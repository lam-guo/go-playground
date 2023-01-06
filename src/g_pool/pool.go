package g_pool

import (
	"sync"
)

type pool struct {
	// slice存放可用worker
	workers []worker
	// 池子大小
	size int
	lock sync.Mutex
	once sync.Once
}

func NewPool(size int) *pool {
	return &pool{
		size:    size,
		lock:    sync.Mutex{},
		workers: nil,
	}
}

func (p *pool) GetWorker() worker {
	p.lock.Lock()

	return worker{}
}

func (p *pool) SetJob(job f) {
	worker := p.GetWorker()
	worker.task <- job
}
