package queue

import (
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/bits-and-blooms/bitset"
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	// GiB // 1073741824
	// TiB // 1099511627776             (超过了int32的范围)
	// PiB // 1125899906842624
	// EiB // 1152921504606846976
	// ZiB // 1180591620717411303424    (超过了int64的范围)
	// YiB // 1208925819614629174706176
)

const (
	TestSize = 10000
	n        = 2000
	loop     = 1
)

// TestMpsc is used to test mpsc
func TestMpsc(t *testing.T) {
	b := bitset.New(n * TestSize)

	curMem := runtime.MemStats{}
	runtime.ReadMemStats(&curMem)
	var wg sync.WaitGroup

	t1 := time.Now()

	for index := 0; index < loop; index++ {

		q := NewMpscQueue()
		for i := 0; i < n; i++ {
			wg.Add(1)
			go func(i int) {
				var v interface{}
				for j := 0; j < TestSize; j++ {
					v = i*TestSize + j
					q.Push(v)
				}
				wg.Done()
			}(i)
		}
		wg.Wait()

		c := 0
		var v interface{}
		for c < n*TestSize {
			v = q.Pop()
			if v != nil {
				b.Set(uint(v.(int)))
				c++
			}

		}
	}

	elapsed := time.Since(t1)
	t.Logf("elapsed: %dns  %dns/op", elapsed, elapsed/(n*TestSize*loop))

	t.Logf("b.count: %d  q.size: %d  b.len: %d", b.Count(), 0, b.Len())

	if !b.All() {
		t.Errorf("All() == false")
	}
	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)

	t.Logf("Alloc = %v MiB", (m.Alloc)/MiB)
	t.Logf("\tTotalAlloc = %v MiB", (m.TotalAlloc)/MiB)
	t.Logf("\tSys = %v MiB", m.Sys/MiB)
	t.Logf("\tNumGC = %v", m.NumGC)
	t.Logf("\tAllocObjCnt = %v", m.Mallocs)
	t.Logf("\tSTW = %v\n", m.PauseTotalNs)
	t.Logf("\tGCCPUFraction = %v\n", m.GCCPUFraction)
}

func TestRingQueue(t *testing.T) {
	b := bitset.New(n * TestSize)

	curMem := runtime.MemStats{}
	runtime.ReadMemStats(&curMem)
	var wg sync.WaitGroup

	t1 := time.Now()
	for index := 0; index < loop; index++ {

		q := NewRingQueue(n * TestSize)
		for i := 0; i < n; i++ {
			wg.Add(1)
			go func(i int) {
				for j := 0; j < TestSize; j++ {
					q.Push(i*TestSize + j)
				}
				wg.Done()
			}(i)
		}
		wg.Wait()

		c := 0
		var v interface{}
		for c < n*TestSize {
			v = q.Pop()
			if v != nil {
				b.Set(uint(v.(int)))
				c++
			}

		}

	}

	elapsed := time.Since(t1)
	t.Logf("elapsed: %dns  %dns/op", elapsed, elapsed/(n*TestSize*loop))

	t.Logf("b.count: %d  q.size: %d  b.len: %d", b.Count(), Roundup2(n*TestSize), b.Len())

	if !b.All() {
		t.Errorf("All() == false")
	}
	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)

	t.Logf("Alloc = %v MiB", (m.Alloc)/MiB)
	t.Logf("\tTotalAlloc = %v MiB", (m.TotalAlloc)/MiB)
	t.Logf("\tSys = %v MiB", m.Sys/MiB)
	t.Logf("\tNumGC = %v", m.NumGC)
	t.Logf("\tAllocObjCnt = %v", m.Mallocs)
	t.Logf("\tSTW = %v\n", m.PauseTotalNs)
	t.Logf("\tGCCPUFraction = %v\n", m.GCCPUFraction)
}
