package trees

import (
	"cmp"
	"fmt"

	"ds-algorithms/pkg/datastructures/array"
	"ds-algorithms/pkg/datastructures/trees/nodes"
)

// TODO --> CREATE BST THAT ALLOWS DUPLICATE KEYS

type BinarySearchTree[T cmp.Ordered] struct {
	root *nodes.BinaryNode[T]
	size int
}

func NewBinarySearchTree[T cmp.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

func (bst *BinarySearchTree[T]) Size() int {
	return bst.size
}

func (bst *BinarySearchTree[T]) IsEmpty() bool {
	return bst.size == 0
}

func (bst *BinarySearchTree[T]) Insert(item T) error {
	var parent *nodes.BinaryNode[T]
	node := bst.root

	for node != nil {
		parent = node

		if item == node.Value {
			return nil
		} else if item < node.Value {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	newNode := &nodes.BinaryNode[T]{Value: item, Left: nil, Right: nil}
	if parent == nil {
		bst.root = newNode
	} else if item < parent.Value {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}

	bst.size++
	return nil
}

func (bst *BinarySearchTree[T]) Search(item T) (T, bool) {
	var zero T
	node := bst.root

	for node != nil {
		if item == node.Value {
			return node.Value, true
		} else if item < node.Value {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return zero, false
}

func (bst *BinarySearchTree[T]) Delete(item T) error {
	initialSize := bst.size
	bst.root = bst.removeHelper(bst.root, item)

	if bst.size < initialSize {
		return nil
	}
	return fmt.Errorf("item %v not found in tree", item)
}

func (bst *BinarySearchTree[T]) TraversePreOrder(node *nodes.BinaryNode[T]) (*array.ArrayList[T], error) {
	var zero *array.ArrayList[T]
	arr := array.NewArrayList[T]()

	err := nodes.PreOrderTraversal(node, arr)
	if err != nil {
		return zero, err
	}

	return arr, nil
}

func (bst *BinarySearchTree[T]) TraversePostOrder(node *nodes.BinaryNode[T]) (*array.ArrayList[T], error) {
	var zero *array.ArrayList[T]
	arr := array.NewArrayList[T]()

	err := nodes.PostOrderTraversal(node, arr)
	if err != nil {
		return zero, err
	}

	return arr, nil
}

func (bst *BinarySearchTree[T]) TraverseInOrder(node *nodes.BinaryNode[T]) (*array.ArrayList[T], error) {
	var zero *array.ArrayList[T]
	arr := array.NewArrayList[T]()

	err := nodes.InOrderTraversal(node, arr)
	if err != nil {
		return zero, err
	}

	return arr, nil
}

func (bst *BinarySearchTree[T]) Clear() error {
	bst.root = nil
	bst.size = 0
	return nil
}

func (bst *BinarySearchTree[T]) largestNode(node *nodes.BinaryNode[T]) *nodes.BinaryNode[T] {
	if node == nil {
		return nil
	}

	currentNode := node
	for currentNode.Right != nil {
		currentNode = currentNode.Right
	}

	return currentNode
}

func (bst *BinarySearchTree[T]) removeHelper(node *nodes.BinaryNode[T], val T) *nodes.BinaryNode[T] {
	if node == nil {
		return nil
	}

	if val < node.Value {
		node.Left = bst.removeHelper(node.Left, val)
		return node
	} else if val > node.Value {
		node.Right = bst.removeHelper(node.Right, val)
		return node
	} else {
		if node.Left == nil {
			bst.size--
			return node.Right
		} else if node.Right == nil {
			bst.size--
			return node.Left
		} else {
			pred := bst.largestNode(node.Left)
			node.Value = pred.Value
			node.Left = bst.removeHelper(node.Left, pred.Value)
			return node
		}
	}
}
