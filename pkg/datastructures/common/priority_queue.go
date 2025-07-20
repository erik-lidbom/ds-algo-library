package common

import (
	"cmp"
)

type MinPriorityQueue[T cmp.Ordered] interface {
	Collection
	Add(elem T)
	RemoveMin() (T, error)
	GetMin() (T, error)
}

type MaxPriorityQueue[T cmp.Ordered] interface {
	Collection
	Add(elem T)
	RemoveMax() (T, error)
	GetMax() (T, error)
}