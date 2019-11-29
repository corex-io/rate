package rate

var mark = struct{}{}

// Limit basic limit
type Limit chan interface{}

// NewLimit create basic limit
func NewLimit(opts ...Option) Limit {
	options := newOptions(opts...)
	return make(chan interface{}, options.Max)
}

// Add add
func (c Limit) Add(delta int) {
	c <- mark
}

// Done done
func (c Limit) Done() {
	<-c
}

// Do execute func with limited, *But* goroutine unsaft
func (c Limit) Do(f func()) {
	c.Add(1)
	go func() {
		defer c.Done()
		f()
	}()
}

// Wait wait
func (c Limit) Wait() {}

// Len len
func (c Limit) Len() int {
	return len(c)
}
