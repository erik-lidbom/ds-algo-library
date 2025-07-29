package linkedlist

import "errors"

type Node struct {
	val  any
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (ll *LinkedList) Size() int {
	return ll.size
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.size == 0
}

func (ll *LinkedList) Add(index int, elem any) error {
	if index < 0 || index > ll.size {
		return errors.New("index out of bounds")
	}

	if index == 0 {
		newNode := &Node{val: elem, next: ll.head}
		ll.head = newNode
		ll.size++
		return nil
	}

	prevNode := ll.head

	for i := 0; i < index-1; i++ {
		prevNode = prevNode.next
	}
	newNode := &Node{val: elem, next: prevNode.next}
	prevNode.next = newNode
	ll.size++

	return nil
}

func (ll *LinkedList) Get(index int) (any, error) {
	if index < 0 || index >= ll.size {
		return nil, errors.New("index out of bounds")
	}

	if index == 0 {
		return ll.head.val, nil
	}

	currentNode := ll.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}
	return currentNode.val, nil
}

func (ll *LinkedList) Set(index int, elem any) error {
	if index < 0 || index >= ll.size {
		return errors.New("index out of bounds")
	}

	if index == 0 {
		ll.head.val = elem
		return nil
	}

	currentNode := ll.head

	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}
	currentNode.val = elem
	return nil
}

func (ll *LinkedList) Remove(index int) (any, error) {
	if index < 0 || index >= ll.size {
		return nil, errors.New("index out of bounds")
	}

	if index == 0 {
		removedNode := ll.head
		ll.head = removedNode.next
		removedNode.next = nil
		ll.size--
		return removedNode.val, nil
	}

	prevNode := ll.head

	for i := 0; i < index-1; i++ {
		prevNode = prevNode.next
	}

	removedNode := prevNode.next
	prevNode.next = removedNode.next
	removedNode.next = nil

	ll.size--
	return removedNode.val, nil
}
