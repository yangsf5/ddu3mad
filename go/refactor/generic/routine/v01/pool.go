package v01

import (
	"sync"
)

type Pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

func NewPool(conCnt, total int) Pool {
	if conCnt < 1 {
		conCnt = 1
	}

	if total < conCnt {
		conCnt = total
	}

	return Pool{
		queue: make(chan int, conCnt),
		wg:    new(sync.WaitGroup),
	}
}

func (m *Pool) AddOne() {
	m.wg.Add(1)
	m.queue <- 1
}

func (m *Pool) DelOne() {
	<-m.queue
	m.wg.Done()
}

func (m *Pool) Waiting() {
	m.wg.Wait()
}
