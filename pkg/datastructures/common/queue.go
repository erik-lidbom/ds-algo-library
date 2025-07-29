package common

type Queue interface {
	Collection
	Enqueue(x any)
	Dequeue() (any, error)
	Peek() (any, error)
}
