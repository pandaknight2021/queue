// implement an efficient bounded queue, multi-producer, single-consumer lock-free queue.

// The Push function is safe to call from multiple goroutines. The Pop and Empty APIs must only be
// called from a single, consumer goroutine.
package queue

import (
	"sync/atomic"
)

type RingQueue struct {
	buf      []interface{}
	head     int64
	tail     int64
	capacity int64
	modbits  int64
	free     int64
}

func NewRingQueue(size int64) *RingQueue {
	cap := int64(Roundup2(uint64(size)))
	return &RingQueue{
		buf:      make([]interface{}, cap),
		head:     0,
		tail:     0,
		modbits:  cap - 1,
		capacity: cap,
		free:     cap,
	}
}

//can be safe to call from multiple goroutines, if q is full , return false
func (q *RingQueue) Push(x interface{}) bool {
	// race the free token
	free := atomic.AddInt64(&q.free, -1) + 1

	if free > 0 {
		// win the race
		tail := atomic.AddInt64(&q.tail, 1)

		q.buf[(tail-1)&q.modbits] = x

		return true
	} else {
		atomic.AddInt64(&q.free, 1)
		return false
	}
}

func (q *RingQueue) Size() int64 {
	return q.capacity - atomic.LoadInt64(&q.free)
}

func (q *RingQueue) Empty() bool {
	return q.Size() == 0
}

//must be call from single consumer
func (q *RingQueue) Pop() interface{} {
	if q.Empty() {
		return nil
	}

	v := q.buf[q.head]
	q.buf[q.head] = nil

	q.head = q.next(q.head)
	atomic.AddInt64(&q.free, 1)

	return v
}

//must be call from single consumer
func (q *RingQueue) Peek() interface{} {
	if q.Empty() {
		return nil
	}

	v := q.buf[q.head]
	return v
}

// prev returns the previous buffer position wrapping around buffer.
func (q *RingQueue) prev(i int64) int64 {
	return (i - 1) & q.modbits // bitwise modulus
}

// next returns the next buffer position wrapping around buffer.
func (q *RingQueue) next(i int64) int64 {
	return (i + 1) & q.modbits // bitwise modulus
}

func (q *RingQueue) resize(size int64) {
	newBuf := make([]interface{}, size)

	l := q.capacity
	//first get the data from q.head until the end of the buffer
	n := min(q.capacity-q.head, l)
	copy(newBuf, q.buf[q.head:(q.head+n)])
	//then get the rest (if any) from the beginning of the buffer
	copy(newBuf[n:], q.buf[:l-n])

	q.buf = newBuf
	q.head = 0
	q.tail = l
	q.modbits = size - 1
	q.capacity = size
}

//roundup to power of 2 for bitwise modulus: x % n == x & (n - 1).
func Roundup2(v uint64) uint64 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v |= v >> 32
	v++
	return v
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}
