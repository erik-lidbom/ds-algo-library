package linkedlist

import "errors"

type LinkedQueue struct {
	front *Node
	rear  *Node
	size  int
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{}
}

func (lq *LinkedQueue) Size() int {
	return lq.size
}

func (lq *LinkedQueue) IsEmpty() bool {
	return lq.size == 0
}

func (lq *LinkedQueue) Enqueue(x any) {
	newNode := &Node{val: x, next: nil}
	if lq.IsEmpty() {
		lq.front = newNode
	} else {
		lq.rear.next = newNode
	}

	lq.rear = newNode
	lq.size++
}

func (lq *LinkedQueue) Dequeue() (any, error) {
	if lq.IsEmpty() {
		return nil, errors.New("cannot dequeue an empty queue")
	}

	dequeuedNode := lq.front
	lq.front = dequeuedNode.next
	dequeuedNode.next = nil

	lq.size--

	if lq.IsEmpty() {
		lq.rear = nil
	}
	return dequeuedNode, nil
}

func (lq *LinkedQueue) Peek() (any, error) {
	if lq.IsEmpty() {
		return nil, errors.New("cannot peek an empty queue")
	}

	return lq.front.val, nil
}
