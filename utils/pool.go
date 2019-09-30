package utils

import (
	"math"
	"sync"
)

// WaitGroup连接池
type WaitGroupPool struct {
	pool chan struct{}
	wg   *sync.WaitGroup
}

// 连接池构造器，入参小于零的话不限制
func NewWaitGroupPool(size int) *WaitGroupPool {
	if size <= 0 {
		size = math.MaxInt32
	}
	return &WaitGroupPool{
		pool: make(chan struct{}, size),
		wg:   &sync.WaitGroup{},
	}
}

// 逐一添加协程实例
func (p *WaitGroupPool) Add() {
	p.pool <- struct{}{}
	p.wg.Add(1)
}

// 完成
func (p *WaitGroupPool) Done() {
	<-p.pool
	p.wg.Done()
}

// 等待连接数
func (p *WaitGroupPool) Wait() {
	p.wg.Wait()
}
