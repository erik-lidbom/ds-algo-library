package common

import (
	"cmp"

	"ds-algorithms/pkg/datastructures/array"
)

type Set[T comparable] interface {
	Collection
	Add(elem T) error
	Remove(elem T) (T, error)
	Contains(elem T) bool
}

type SortedSet[T cmp.Ordered] interface {
	Set[T]
	First() (T, error)
	Last() (T, error)
	Floor(x T) (T, error)
	Ceiling(x T) (T, error)
	Lower(x T) (T, error)
	Higher(x T) (T, error)
	Between(x, y T) *array.ArrayList[T]
}
