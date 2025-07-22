package sets

import (
	"cmp"
	"fmt"
)

type SetNode[T any] struct {
	elem T
	next *SetNode[T]
}

type LinkedSet[T cmp.Ordered] struct {
	head *SetNode[T]
	size int
}

func NewLinkedSet[T cmp.Ordered]() *LinkedSet[T] {
	return &LinkedSet[T]{head: nil, size: 0}
}

func (ls *LinkedSet[T]) Size() int {
	return ls.size
}

func (ls *LinkedSet[T]) IsEmpty() bool {
	return ls.size == 0
}

func (ls *LinkedSet[T]) Add(elem T) error {
	containVal := ls.Contains(elem)

	if containVal {
		return fmt.Errorf("cannot add element to set: value %v already exists\n", containVal)
	}

	newHead := &SetNode[T]{elem: elem, next: ls.head}
	ls.head = newHead
	ls.size++
	return nil
}

func (ls *LinkedSet[T]) Remove(elem T) (T, error) {
	var zero T
	var prevNode *SetNode[T]
	currNode := ls.head

	for currNode != nil {
		if elem == currNode.elem {
			if prevNode == nil {
				ls.head = currNode.next
			} else {
				prevNode.next = currNode.next
			}
			currNode.next = nil
			ls.size--
			return  elem, nil
		}
	}
	return zero, fmt.Errorf("element %v not found in the set", elem)
}

func (ls *LinkedSet[T]) Contains(elem T) bool {
	node := ls.head

	for node != nil {
		if elem == node.elem {
			return true
		}
		node = node.next
	}
	return false
}