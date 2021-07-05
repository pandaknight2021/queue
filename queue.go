package queue

// Queue is a FIFO data structure.
// Push puts a value into its tail,
// Pop removes a value from its head.
type Queue interface {
	Push(v interface{}) bool
	Pop() interface{}
	Empty() bool
	Size() int64
}
