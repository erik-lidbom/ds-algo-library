package searchable

import (
	"cmp"
	"errors"
)


// SearchableSlice works as a wrapper that can be used on normal Go arrays and slices. 
// Since the algorithms provided in this project must work across different data structures, such as Arraylist,
// we need a common interface that both standard slices and custom structures can implement.

var ErrIndexOutOfBounds = errors.New("index out of bounds")

type SearchableSlice[T cmp.Ordered] []T

func (ss SearchableSlice[T]) Size() int {
	return len(ss)
}

func (ss SearchableSlice[T]) Get(index int) (T, error) {
	if index < 0 || index >= len(ss) {
		var zeroVal T
		return zeroVal, ErrIndexOutOfBounds
	}
	return ss[index], nil
}