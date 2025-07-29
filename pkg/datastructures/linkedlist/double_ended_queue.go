package linkedlist

import "errors"

type DoubleNode struct {
	val  any
	prev *DoubleNode
	next *DoubleNode
}

type DoubleDeque struct {
	head *DoubleNode
	tail *DoubleNode
	size int
}

func (dd *DoubleDeque) Size() int {
	return dd.size
}

func (dd *DoubleDeque) IsEmpty() bool {
	return dd.size == 0
}

func (dd *DoubleDeque) AddFirst(elem any) {
	if dd.size == 0 {
		newNode := &DoubleNode{val: elem, prev: nil, next: nil}
		dd.head = newNode
		dd.tail = newNode
	} else {
		newNode := &DoubleNode{val: elem, prev: nil, next: dd.head}
		dd.head.prev = newNode
		dd.head = newNode
	}
	dd.size++
}

func (dd *DoubleDeque) AddLast(elem any) {
	if dd.size == 0 {
		newNode := &DoubleNode{val: elem, prev: nil, next: nil}
		dd.head = newNode
		dd.tail = newNode
	} else {
		newNode := &DoubleNode{val: elem, prev: dd.tail, next: nil}
		dd.tail.next = newNode
		dd.tail = newNode
	}
	dd.size++
}

func (dd *DoubleDeque) RemoveFirst() (any, error) {
	if dd.IsEmpty() {
		return nil, errors.New("cannot remove element from empty list")
	}

	removedElem := dd.head
	dd.head = removedElem.next

	if dd.head != nil {
		dd.head.prev = nil
	}

	if dd.size == 1 {
		dd.tail = nil
	}

	removedElem.next = nil
	dd.size--
	return removedElem.val, nil
}

func (dd *DoubleDeque) RemoveLast() (any, error) {
	if dd.IsEmpty() {
		return nil, errors.New("cannot remove element from empty list")
	}

	removedElem := dd.tail
	dd.tail = removedElem.prev

	if dd.tail != nil {
		dd.tail.next = nil
	}

	if dd.size == 1 {
		dd.head = nil
	}

	removedElem.prev = nil
	dd.size--
	return removedElem.val, nil
}

func (dd *DoubleDeque) PeekFirst() (any, error) {
	if dd.IsEmpty() {
		return nil, errors.New("cannot peek empty list")
	}
	return dd.head.val, nil
}

func (dd *DoubleDeque) PeekLast() (any, error) {
	if dd.IsEmpty() {
		return nil, errors.New("cannot peek empty list")
	}
	return dd.tail.val, nil
}
