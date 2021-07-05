# golang-queue
queue provides an efficient implementation of a multi-producer, single-consumer lock-free queue
 
Reference

MpscQueue copy from :  https://github.com/AsynkronIT/protoactor-go

Interface:
type Queue interface {
	Push(v interface{}) bool
	Pop() interface{}
	Empty() bool
	Size() int64
}

usage: 

 q := NewRingQueue(size)
 q.Push(value)
 
 //note: pop only be call from single gorotine
 v = q.Pop()
 
benchmark test:

test env :  i5 8th GEN / 16G
(2000 goroutine * 10000) push  && pop

goos: linux
goarch: amd64
pkg: queue
BenchmarkRingqueue
BenchmarkRingqueue-8      	       1	2645369800 ns/op	  853040 B/op	    1443 allocs/op
BenchmarkRingqueueBig
BenchmarkRingqueueBig-8   	       1	8381621100 ns/op	537102736 B/op	     561 allocs/op
BenchmarkMpscqueue
BenchmarkMpscqueue-8      	       1	8437661000 ns/op	640468328 B/op	20003457 allocs/op
BenchmarkCQueue
BenchmarkCQueue-8         	       1	13675927000 ns/op	640196208 B/op	20002046 allocs/op
PASS
coverage: 67.0% of statements
ok  	queue	33.540s

