package array

import (
	"errors"
)


type ArrayStack struct {
	arr []any
	size int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{arr: make([]any, 1), size: 0}
}

func (as *ArrayStack) Size() int {
	return as.size
}

func (as *ArrayStack) IsEmpty() bool {
	return as.size == 0
}

func (as *ArrayStack) Push(x any) {
	if as.size >= len(as.arr) {
		as.resizeArray()
	}
	arraySize := as.size

	as.arr[arraySize] = x
	as.size++
}

func (as *ArrayStack) Pop() (any, error) {

	if as.IsEmpty() {
		return nil, errors.New("cannot pop from an empty stack")
	}

	arraySize := as.size
	removedVal := as.arr[arraySize - 1]
	as.arr[arraySize - 1] = nil
	as.size--

	if as.size * 3 <= len(as.arr) {
		as.shrinkArray()
	}
	return removedVal, nil
}

func (as *ArrayStack) Peek() (any, error) {

	if as.IsEmpty() {
		return nil, errors.New("cannot peek an empty stack")
	}
	arraySize := as.size
	peekValue := as.arr[arraySize - 1]
	return peekValue, nil
}

func (as *ArrayStack) resizeArray() {
	oldArray := as.arr
	as.arr = make([]any, as.size * 2)

	for i := 0; i < as.size; i++ {
		as.arr[i] = oldArray[i]
	}
}

func (as *ArrayStack) shrinkArray() {
	oldArray := as.arr
	as.arr = make([]any, len(as.arr) / 2)

	for i := 0; i < as.size; i++ {
		as.arr[i] = oldArray[i]
	}
}