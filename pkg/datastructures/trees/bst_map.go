package trees

import (
	"cmp"
	"fmt"

	"ds-algorithms/pkg/datastructures/array"
	"ds-algorithms/pkg/datastructures/trees/nodes"
)

type BinarySearchTreeMap[K cmp.Ordered, V comparable] struct {
	root *nodes.BinaryMapNode[K, V]
	size int
}

func NewBinarySearchTreeMap[K cmp.Ordered, V comparable]() *BinarySearchTreeMap[K, V] {
	return &BinarySearchTreeMap[K, V]{
		root: nil,
		size: 0,
	}
}

func (bstm *BinarySearchTreeMap[K, V]) Size() int {
	return bstm.size
}

func (bstm *BinarySearchTreeMap[K, V]) IsEmpty() bool {
	return bstm.size == 0
}

func (bstm *BinarySearchTreeMap[K, V]) Put(key K, value V) {
	bstm.putHelper(bstm.root, key, value)
}

func (bstm *BinarySearchTreeMap[K, V]) Get(key K) (V, error) {
	var zero V
	node := bstm.getHelper(bstm.root, key)

	if node == nil {
		return zero, fmt.Errorf("cannot find key in treemap")
	}

	return node.Value, nil
}

func (bstm *BinarySearchTreeMap[K, V]) Remove(key K) (V, error) {
	var zeroV V

	newRoot, removedValue, wasRemoved := bstm.removeHelper(bstm.root, key)
	bstm.root = newRoot

	if wasRemoved {
		bstm.size--
		return removedValue, nil
	}

	return zeroV, fmt.Errorf("item %v not found in tree", key)
}

func (bstm *BinarySearchTreeMap[K, V]) ContainsKey(key K) bool {
	node := bstm.getHelper(bstm.root, key)
	if node == nil {
		return false
	}
	return true
}

func (bstm *BinarySearchTreeMap[K, V]) Insert(key K) error {
	var zeroVal V
	bstm.Put(key, zeroVal)
	return nil
}

func (bstm *BinarySearchTreeMap[K, V]) Search(key K) (V, bool) {
	var zero V
	nodeVal, err := bstm.Get(key)
	if err != nil {
		return zero, false
	}

	return nodeVal, true
}

func (bstm *BinarySearchTreeMap[K, V]) Delete(key K) error {
	initialSize := bstm.size
	bstm.root, _, _ = bstm.removeHelper(bstm.root, key)

	if bstm.size < initialSize {
		return nil
	}

	return fmt.Errorf("key %v not found in tree", key)
}

func (bstm *BinarySearchTreeMap[K, V]) KeysBetween(min, max K) *array.ArrayList[K] {
	result := array.NewArrayList[K]()

	if bstm.IsEmpty() || min > max {
		return result
	}
}

func (bstm *BinarySearchTreeMap[K, V]) TraversePreOrder(node *nodes.BinaryMapNode[K, V]) (*array.ArrayList[K], error) {
	var zero *array.ArrayList[K]
	arr := array.NewArrayList[K]()

	err := nodes.PreOrderTraversal(node, arr)
	if err != nil {
		return zero, err
	}

	return arr, nil
}

func (bstm *BinarySearchTreeMap[K, V]) TraversePostOrder(node *nodes.BinaryMapNode[K, V]) (*array.ArrayList[K], error) {
	var zero *array.ArrayList[K]
	arr := array.NewArrayList[K]()

	err := nodes.PostOrderTraversal(node, arr)
	if err != nil {
		return zero, err
	}

	return arr, nil
}

func (bstm *BinarySearchTreeMap[K, V]) TraverseInOrder(node *nodes.BinaryMapNode[K, V]) (*array.ArrayList[K], error) {
	var zero *array.ArrayList[K]
	arr := array.NewArrayList[K]()

	err := nodes.InOrderTraversal(node, arr)
	if err != nil {
		return zero, err
	}

	return arr, nil
}

func (bstm *BinarySearchTreeMap[K, V]) largestNode(node *nodes.BinaryMapNode[K, V]) *nodes.BinaryMapNode[K, V] {
	if node == nil {
		return nil
	}

	currentNode := node
	for currentNode.Right != nil {
		currentNode = currentNode.Right
	}

	return currentNode
}

func (bstm *BinarySearchTreeMap[K, V]) getHelper(node *nodes.BinaryMapNode[K, V], key K) *nodes.BinaryMapNode[K, V] {
	if node == nil {
		return nil
	} else if key < node.Key {
		return bstm.getHelper(node.Left, key)
	} else if key > node.Key {
		return bstm.getHelper(node.Right, key)
	} else {
		return node
	}
}

func (bstm *BinarySearchTreeMap[K, V]) putHelper(node *nodes.BinaryMapNode[K, V], key K, value V) *nodes.BinaryMapNode[K, V] {
	if node == nil {
		bstm.size++
		return &nodes.BinaryMapNode[K, V]{
			Key:   key,
			Value: value,
			Left:  nil,
			Right: nil,
		}
	} else if key < node.Key {
		bstm.putHelper(node.Left, key, value)
	} else if key > node.Key {
		bstm.putHelper(node.Right, key, value)
	}
	node.Value = value
	return node
}

func (bstm *BinarySearchTreeMap[K, V]) removeHelper(node *nodes.BinaryMapNode[K, V], key K) (*nodes.BinaryMapNode[K, V], V, bool) {
	var zeroV V

	if node == nil {
		return nil, zeroV, false
	}

	var removedValue V
	var removed bool

	if key < node.Key {
		node.Left, removedValue, removed = bstm.removeHelper(node.Left, key)
		return node, removedValue, removed
	} else if key > node.Key {
		node.Right, removedValue, removed = bstm.removeHelper(node.Right, key)
		return node, removedValue, removed
	} else {
		if node.Left == nil {
			return node.Right, node.Value, true
		} else if node.Right == nil {
			return node.Left, node.Value, true
		} else {
			pred := bstm.largestNode(node.Left)
			originalNodeValue := node.Value

			node.Key = pred.Key
			node.Value = pred.Value

			node.Left, _, _ = bstm.removeHelper(node.Left, pred.Key)

			return node, originalNodeValue, true
		}
	}
}

func (bstm *BinarySearchTreeMap[K, V]) keysBetweenHelper(node *nodes.BinaryMapNode[K, V], min K, max K, result *array.ArrayList[K]) {
	if node == nil {
		return
	}

	if min < node.Key {
		bstm.keysBetweenHelper(node.Left, min, max, result)
	}

	if min <= node.Key && node.Key <= max {
		result.Add(result.Size(), node.Key)
	}

	if node.Key < max {
		bstm.keysBetweenHelper(node.Right, min, max, result)
	}
}
