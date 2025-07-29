package common

import (
	"cmp"
)

// Searchable represents a data structure that allows elements to be retrieved by index and provides its total size
// Size() - returns the amount of elements in the slice, ArrayList or other Searchable compatible types
// Get() - returns an element of type T and an error
// Set() - sets a value in a Searchable list without increasing the size of the list
type Searchable[T cmp.Ordered] interface {
	Size() int
	Get(index int) (T, error)
	Set(index int, elem T) error
}
