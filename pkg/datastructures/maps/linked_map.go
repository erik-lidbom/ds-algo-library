package maps

import (
	"cmp"
	"fmt"
)

type MapNode[K cmp.Ordered, V any] struct {
	key K
	value V
	next *MapNode[K, V]
}

type LinkedMap[K cmp.Ordered, V any] struct {
	head *MapNode[K, V]
	size int
}

func NewLinkedMap[K cmp.Ordered, V any]() *LinkedMap[K, V] {
	return &LinkedMap[K, V]{head: nil, size: 0}
}

func (lm *LinkedMap[K, V]) Size() int {
	return lm.size
}

func (lm *LinkedMap[K, V]) IsEmpty() bool {
	return lm.size == 0
}

func (lm *LinkedMap[K, V]) Put(key K, value V)  {
	currNode := lm.head

	for currNode != nil {
		if key == currNode.key {
			currNode.value = value
			return 
		}
		currNode = currNode.next
	}
	newNode := &MapNode[K, V]{key: key, value: value, next: lm.head}
	lm.head = newNode
	lm.size++
}

func (lm *LinkedMap[K, V]) Get(key K) (V, error) {
	var zero V

	node := lm.head

	for node != nil {
		if key == node.key {
			return node.value, nil
		}
		node = node.next
	}
	return zero, fmt.Errorf("cannot retrieve value: key '%v' not found in the map", key)
}

func (lm *LinkedMap[K, V]) Remove(key K) (V, error) {
	var zero V
	var prevNode *MapNode[K, V]
	currNode := lm.head

	for currNode != nil {
		if key == currNode.key {
			removedVal := currNode.value
			if prevNode == nil {
				lm.head = currNode.next
			} else {
				prevNode.next = currNode.next
			}
			currNode.next = nil
			lm.size--
			return removedVal, nil
		}
		prevNode = currNode
		currNode = currNode.next
	}
	return zero, fmt.Errorf("key %v not found in the map", key)
}

func (lm *LinkedMap[K, V]) ContainsKey(key K) bool {
	node := lm.head

	for node != nil {
		if key == node.key {
			return true
		}
		node = node.next
	}
	return false
}