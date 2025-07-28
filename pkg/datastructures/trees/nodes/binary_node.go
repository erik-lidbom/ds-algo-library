package nodes

import "cmp"

type BinaryNode[T cmp.Ordered] struct {
	Value T
	Left *BinaryNode[T]
	Right *BinaryNode[T]
}

func (bn *BinaryNode[T]) GetLeft() TraversableNode[T] {
	if bn == nil || bn.Left == nil {
		return nil
	}
	return bn.Left
}

func (bn *BinaryNode[T]) GetRight() TraversableNode[T] {
	if bn == nil || bn.Right == nil {
		return nil
	}
	return bn.Right
}

func (bn *BinaryNode[T]) GetValue() T {
	if bn == nil {
		var zero T
		return zero
	}
	return bn.Value
}

func (bn *BinaryNode[T]) IsNil() bool {
	return bn == nil
}