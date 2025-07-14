package common

type Deque interface {
	Collection
	AddFirst(x any)
	AddLast(a any)
	RemoveFirst() (any, error)
	RemoveLast() (any, error)
	PeekFirst() (any, error)
	PeekLast() (any, error)
}