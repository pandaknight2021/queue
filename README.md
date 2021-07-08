# golang-queue
queue provides an efficient implementation of a multi-producer, single-consumer lock-free queue

## Reference

-  [MpscQueue] (github.com/AsynkronIT/protoactor-go)

## Interface

``` go
type Queue interface {
    Push(v interface{}) bool
    Pop() interface{}
    Empty() bool
    Size() int64
}
```

## usage

``` go
q := NewRingQueue(size)
q.Push(value)

//note: pop only be call from single goroutine
v = q.Pop()
```

## benchmark

test env :  i5 8th GEN / 16G  
2000 goroutine * 10000 push  && pop  

**RingQueue:  420ns/op**  
**MpscQueue:  640ns/op**  
**CQueue:     830ns/op**  

``` console
goos: linux
goarch: amd64
pkg: queue
BenchmarkRingqueue
BenchmarkRingqueue-8   	       1	8396128600 ns/op	697504768 B/op	20001304 allocs/op
BenchmarkMpscqueue
BenchmarkMpscqueue-8   	       1	12918583100 ns/op	800511184 B/op	40002378 allocs/op
BenchmarkCQueue
BenchmarkCQueue-8      	       1	16738020500 ns/op	800215496 B/op	40001946 allocs/op
PASS
coverage: 64.9% of statements
ok  	queue	38.523s
```

