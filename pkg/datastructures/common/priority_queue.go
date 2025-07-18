package common

type MinPriorityQueue interface {
	Collection
	Add(elem any)
	RemoveMin() any
	GetMin() any
}

type MaxPriorityQueue interface {
	Collection
	Add(elem any)
	RemoveMax() any
	GetMax() any
}