package linkedlist

import "errors"

type SinglyLinkedList struct {
	head *Node
	size int
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (sl *SinglyLinkedList) Size() int {
	return sl.size
}

func (sl *SinglyLinkedList) IsEmpty() bool {
	return sl.size == 0
}

func (sl *SinglyLinkedList) Push(x any) {
	new_node := &Node{
		val:  x,
		next: sl.head,
	}

	sl.head = new_node
	sl.size++
}

func (sl *SinglyLinkedList) Pop() (any, error) {
	if sl.IsEmpty() {
		return nil, errors.New("cannot pop an element of an empty list")
	}
	popped_node := sl.head
	sl.head = sl.head.next
	sl.size--

	return popped_node.val, nil
}

func (sl *SinglyLinkedList) Peek() (any, error) {
	if sl.IsEmpty() {
		return nil, errors.New("cannot peek an empty list")
	}

	return sl.head.val, nil
}
