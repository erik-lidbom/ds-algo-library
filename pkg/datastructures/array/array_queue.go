package array

import (
	"errors"
)

type ArrayQueue struct {
	arr   []any
	front int
	rear  int
	size  int
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{
		arr:   make([]any, 10),
		front: 0,
		rear:  0,
		size:  0,
	}
}

func (aq *ArrayQueue) Size() int {
	return aq.size
}

func (aq *ArrayQueue) IsEmpty() bool {
	return aq.size == 0
}

func (aq *ArrayQueue) Enqueue(item any) error {
	if aq.size >= len(aq.arr) {
		aq.resizeArray()
	}

	aq.arr[aq.rear] = item
	aq.rear = (aq.rear + 1) % len(aq.arr)
	aq.size++
	return nil
}

func (aq *ArrayQueue) Dequeue() (any, error) {
	if aq.IsEmpty() {
		return nil, errors.New("cannot dequeue an empty queue")
	}

	removedItem := aq.arr[aq.front]
	aq.arr[aq.front] = nil
	aq.front = (aq.front + 1) % len(aq.arr)
	aq.size--

	if aq.size*3 <= len(aq.arr) {
		aq.shrinkArray()
	}
	return removedItem, nil
}

func (aq *ArrayQueue) Peek() (any, error) {
	if aq.IsEmpty() {
		return nil, errors.New("cannot peek an empty queue")
	}

	return aq.arr[aq.front], nil
}

func (aq *ArrayQueue) resizeArray() {
	oldArray := aq.arr
	aq.arr = make([]any, aq.size*2)

	for i := 0; i < aq.size; i++ {
		newPos := (aq.front + i) % len(aq.arr)
		aq.arr[i] = oldArray[newPos]
	}
	aq.front = 0
	aq.rear = aq.size
}

func (aq *ArrayQueue) shrinkArray() {
	oldArray := aq.arr
	aq.arr = make([]any, len(aq.arr)/2)

	for i := 0; i < aq.size; i++ {
		newPos := (aq.front + i) % len(oldArray)
		aq.arr[i] = oldArray[newPos]
	}

	aq.front = 0
	aq.rear = aq.size
}
