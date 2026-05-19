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

// NewPool creates a goroutine pool, size: concurrency count, errorBuffer: error buffer size
func NewPool(size, errorBuffer int) *GoroutinePool {
	return &GoroutinePool{
		taskChan:  make(chan Task, size*2),
		sem:       make(chan struct{}, size),
		stopChan:  make(chan struct{}),
		errorChan: make(chan error, errorBuffer),
	}
}

// Submit submits a task (with timeout mechanism)
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

// Run starts worker goroutines
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

// Shutdown gracefully shuts down
func (p *GoroutinePool) Shutdown() {
	close(p.stopChan)
	p.wg.Wait()
	close(p.errorChan)
}

// Errors returns the error channel
func (p *GoroutinePool) Errors() <-chan error {
	return p.errorChan
}

// Usage example
func main() {
	pool := NewPool(10, 100)
	go pool.Run()

	// Submit task
	for i := 0; i < 100; i++ {
		id := i
		pool.Submit(func() error {
			println("processing task", id)
			time.Sleep(100 * time.Millisecond)
			return nil
		}, time.Second)
	}

	// Handle errors
	go func() {
		for err := range pool.Errors() {
			println("task error:", err.Error())
		}
	}()

	time.Sleep(5 * time.Second)
	pool.Shutdown()
}
