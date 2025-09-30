package main

import (
	"sync"
	"time"
)

type Task func() error

type GoroutinePool struct {
	taskChan  chan Task
	wg        sync.WaitGroup
	sem       chan struct{}
	stopChan  chan struct{}
	errorChan chan error
}

// NewPool 创建协程池 size: 并发数量 errorBuffer: 错误缓冲大小
func NewPool(size, errorBuffer int) *GoroutinePool {
	return &GoroutinePool{
		taskChan:  make(chan Task, size*2),
		sem:       make(chan struct{}, size),
		stopChan:  make(chan struct{}),
		errorChan: make(chan error, errorBuffer),
	}
}

// Submit 提交任务（带超时机制）
func (p *GoroutinePool) Submit(task Task, timeout time.Duration) bool {
	select {
	case p.taskChan <- task:
		return true
	case <-time.After(timeout):
		return false
	case <-p.stopChan:
		return false
	}
}

// Run 启动工作协程
func (p *GoroutinePool) Run() {
	for {
		select {
		case task := <-p.taskChan:
			p.sem <- struct{}{}
			p.wg.Add(1)
			go func(t Task) {
				defer func() {
					<-p.sem
					p.wg.Done()
				}()
				if err := t(); err != nil {
					select {
					case p.errorChan <- err:
					default:
					}
				}
			}(task)
		case <-p.stopChan:
			return
		}
	}
}

// Shutdown 优雅关闭
func (p *GoroutinePool) Shutdown() {
	close(p.stopChan)
	p.wg.Wait()
	close(p.errorChan)
}

// Errors 获取错误通道
func (p *GoroutinePool) Errors() <-chan error {
	return p.errorChan
}

// 使用示例
func main() {
	pool := NewPool(10, 100)
	go pool.Run()

	// 提交任务
	for i := 0; i < 100; i++ {
		id := i
		pool.Submit(func() error {
			println("processing task", id)
			time.Sleep(100 * time.Millisecond)
			return nil
		}, time.Second)
	}

	// 处理错误
	go func() {
		for err := range pool.Errors() {
			println("task error:", err.Error())
		}
	}()

	time.Sleep(5 * time.Second)
	pool.Shutdown()
}
