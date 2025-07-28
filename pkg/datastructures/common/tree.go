package common

import (
	"cmp"
	"ds-algorithms/pkg/datastructures/array"
)

type Tree[T cmp.Ordered] interface {
	Collection
	Insert(item T) error
	Search(item T) (T, bool)
	Delete(item T) error
	TraversePreOrder() (array.ArrayList[T], error)
	TraversePostOrder() (array.ArrayList[T], error)
	TraverseInOrder() (array.ArrayList[T], error)
	Clear() error
}