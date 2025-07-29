package searchable

import (
	"cmp"
	"ds-algorithms/pkg/datastructures/common"
	"errors"
	"fmt"
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
		var zero T
		return zero, ErrIndexOutOfBounds
	}
	return ss[index], nil
}

func (ss SearchableSlice[T]) Set(index int, elem T) error {
	if index < 0 || index >= len(ss) {
		return ErrIndexOutOfBounds
	}
	ss[index] = elem
	return nil 
}

func Swap[T cmp.Ordered](arr common.Searchable[T], i, j int) error {
	iValue, err := arr.Get(i)
	if err != nil {
		return fmt.Errorf("failed to retrieve element for index %d\nerror: %w", i, err)
	}

	jValue, err := arr.Get(j)
	if err != nil {
		return fmt.Errorf("failed to retrieve element for index %d\nerror: %w", j, err)
	}

	err = arr.Set(j, iValue)
	if err != nil {
		return fmt.Errorf("failed to swap element: %w", err)
	}

	err = arr.Set(i, jValue) 
	if err != nil {
		// Since the first swap worked as expected, we need to do a rollback.
		rollbackErr := arr.Set(j, jValue)
		if rollbackErr != nil {
			return fmt.Errorf("critical swap error: failed to set element at index %d (original error: %w), AND rollback for index %d failed: %w", j, err, i, rollbackErr)
		}
		return fmt.Errorf("failed to swap element: %w", err) 
	}

	return nil
}