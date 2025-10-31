package workerpoolUtils

import (
	"context"
	"sync"
	"sync/atomic"
)

type Task func()

type Pool struct {
	capacity int32 // 最大并发度
	core     int32 // 最小并发度
	running  int32 // 当前运行 goroutine 数
	taskCh   chan Task
	wg       sync.WaitGroup
	cancel   context.CancelFunc
}

// New 创建动态池
func New(core, capacity int) *Pool {
	if core <= 0 || capacity <= 0 || core > capacity {
		panic("invalid core/capacity")
	}
	ctx, cancel := context.WithCancel(context.Background())
	p := &Pool{
		capacity: int32(capacity),
		core:     int32(core),
		taskCh:   make(chan Task),
		cancel:   cancel,
	}
	// 预先启动 core 个 worker
	for i := 0; i < core; i++ {
		p.spawn(ctx)
	}
	return p
}

// 启动一个 worker
func (p *Pool) spawn(ctx context.Context) {
	p.wg.Add(1)
	atomic.AddInt32(&p.running, 1)
	go func() {
		defer func() {
			atomic.AddInt32(&p.running, -1)
			p.wg.Done()
		}()
		for {
			select {
			case <-ctx.Done():
				return
			case t := <-p.taskCh:
				if t != nil {
					t()
				}
			}
		}
	}()
}

// Submit 提交任务，必要时扩容
func (p *Pool) Submit(t Task) {
	select {
	case p.taskCh <- t:
	default:
		// 队列满，尝试扩容
		if atomic.LoadInt32(&p.running) < atomic.LoadInt32(&p.capacity) {
			p.spawn(context.TODO())
		}
		p.taskCh <- t
	}
}

// Stop 平滑关闭
func (p *Pool) Stop() {
	close(p.taskCh)
	p.cancel()
	p.wg.Wait()
}
