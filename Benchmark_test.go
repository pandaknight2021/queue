package queue

import (
	"sync"
	"testing"
	"time"
)

const (
	RunTimes           = 1000000
	BenchParam         = 10
	BenchAntsSize      = 2000
	DefaultExpiredTime = 10 * time.Second
)

func demoFunc() {
	// num := 10
	// strconv.Itoa(num)
	time.Sleep(time.Duration(BenchParam) * time.Millisecond)

}

func BenchmarkRingqueue(b *testing.B) {
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		q := NewRingQueue(BenchAntsSize)
		for k := 0; k < 10; k++ {
			wg.Add(1)
			go func() {
				for j := 0; j < RunTimes; j++ {
					q.Push(demoFunc)
				}
				wg.Done()
			}()
		}

		for j := 0; j <= RunTimes*10+1; j++ {
			q.Pop()
		}
		wg.Wait()
	}
}

func BenchmarkRingqueueBig(b *testing.B) {
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		q := NewRingQueue(10 * RunTimes)
		for k := 0; k < 10; k++ {
			wg.Add(1)
			go func() {
				for j := 0; j < RunTimes; j++ {
					q.Push(demoFunc)
				}
				wg.Done()
			}()
		}

		for j := 0; j <= RunTimes*10+1; j++ {
			q.Pop()
		}
		wg.Wait()
	}
}

func BenchmarkMpscqueue(b *testing.B) {
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		q := NewMpscQueue()
		for k := 0; k < 10; k++ {
			wg.Add(1)
			go func() {
				for j := 0; j < RunTimes; j++ {
					q.Push(demoFunc)
				}
				wg.Done()
			}()
		}

		for j := 0; j <= RunTimes*10+1; j++ {
			q.Pop()
		}
		wg.Wait()
	}
}

func BenchmarkCQueue(b *testing.B) {
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		q := NewCQueue()
		for k := 0; k < 10; k++ {
			wg.Add(1)
			go func() {
				for j := 0; j < RunTimes; j++ {
					q.Push(demoFunc)
				}
				wg.Done()
			}()
		}

		for j := 0; j <= RunTimes*10+1; j++ {
			q.Pop()
		}
		wg.Wait()
	}
}
