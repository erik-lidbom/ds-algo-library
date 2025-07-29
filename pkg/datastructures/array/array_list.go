package array

import (
	"errors"
	"fmt"
	"strings"
)

type ArrayList[T comparable] struct {
	arr  []T
	size int
}

func NewArrayList[T comparable]() *ArrayList[T] {
	return &ArrayList[T]{arr: make([]T, 10), size: 0}
}

func (al *ArrayList[T]) Size() int {
	return al.size
}

func (al *ArrayList[T]) IsEmpty() bool {
	return al.size == 0
}

func (al *ArrayList[T]) Add(index int, elem T) error {
	if index < 0 || index > al.size {
		return errors.New("index out of bounds")
	}

	if al.size >= len(al.arr) {
		al.resize()
	}

	if index == al.size {
		al.arr[index] = elem
		al.size++
		return nil
	}

	for i := al.size; i > index; i-- {
		al.arr[i] = al.arr[i-1]
	}
	al.arr[index] = elem
	al.size++

	return nil
}

func (al *ArrayList[T]) Get(index int) (T, error) {
	if index < 0 || index >= al.size {
		var zero T
		return zero, errors.New("index out of bounds")
	}

	return al.arr[index], nil
}

func (al *ArrayList[T]) Set(index int, elem T) error {
	if index < 0 || index >= al.size {
		return errors.New("index out of bounds")
	}

	al.arr[index] = elem
	return nil
}

func (al *ArrayList[T]) Remove(index int) (T, error) {
	var zero T

	if index < 0 || index >= al.size {
		return zero, errors.New("index out of bounds")
	}

	removedVal := al.arr[index]

	for i := index + 1; i < al.size; i++ {
		al.arr[i-1] = al.arr[i]
	}
	al.arr[al.size-1] = zero
	al.size--

	if al.size*3 <= len(al.arr) {
		al.shrink()
	}
	return removedVal, nil
}

func (al *ArrayList[T]) String() string {
	if al.size == 0 {
		return "ArrayList: [] (Size: 0)"
	}

	var sb strings.Builder
	sb.WriteString("ArrayList: [")
	for i := 0; i < al.size; i++ {
		sb.WriteString(fmt.Sprintf("%v", al.arr[i]))
		if i < al.size-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(fmt.Sprintf("] (Size: %d)", al.size))
	return sb.String()
}

func (al *ArrayList[T]) resize() {
	oldArray := al.arr

	al.arr = make([]T, len(al.arr)*2)
	for i := 0; i < al.size; i++ {
		al.arr[i] = oldArray[i]
	}
}

func (al *ArrayList[T]) shrink() {
	oldArray := al.arr
	al.arr = make([]T, len(al.arr)/2)

	for i := 0; i < al.size; i++ {
		al.arr[i] = oldArray[i]
	}
}

func Swap[T comparable](arr *ArrayList[T], i, j int) error {
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
