package rate_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/corex-io/rate"
)

type loop int

func (i loop) Run() {
	time.Sleep(1 * time.Second)
	fmt.Println(i, time.Now().Format("2006/01/02 15:04:05.000"))
}

func Test_Limit(t *testing.T) {
	r := rate.NewWaitGroup(rate.Max(3))
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println(r.Len())
		}
	}()
	for {
		// time.Sleep(2*time.Second)
		r.Do(loop(1).Run)
	}
}

func Benchmark_Limit(b *testing.B) {
	l := rate.NewWaitGroup(rate.Max(5000000000000))
	for i := 0; i < b.N; i++ {
		l.Do(loop(i).Run)
	}
	l.Wait()
}
