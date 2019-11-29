package rate

import (
	"sync"
)

// WaitGroup 并发阻塞, 安全退出
type WaitGroup struct {
	Limit
	opts Options
	wg sync.WaitGroup
}

// NewWaitGroup create WaitGroup
func NewWaitGroup(opts ...Option) *WaitGroup {
	return &WaitGroup{
		Limit: NewLimit(opts...),
	}
}

// Add add
func (l *WaitGroup) Add(delta int) {
	l.Limit.Add(delta)
	l.wg.Add(delta)
}

// Done done
func (l *WaitGroup) Done() {
	l.Limit.Done()
	l.wg.Done()
}

// Wait wait
func (l *WaitGroup) Wait() {
	l.wg.Wait()
	close(l.Limit)
}

// Do execute func with limited, *But* goroutine unsaft
func (l *WaitGroup)  Do(f func()) {
	l.Add(1)
	go func() {
		defer l.Done()
		f()
	}()
}

// Close close
func (l *WaitGroup) Close() error {
	close(l.Limit)
	return nil
}
