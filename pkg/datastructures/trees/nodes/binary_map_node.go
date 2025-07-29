package nodes

import "cmp"

type BinaryMapNode[K cmp.Ordered, V comparable] struct {
	Key   K
	Value V
	Left  *BinaryMapNode[K, V]
	Right *BinaryMapNode[K, V]
}

func (bmn *BinaryMapNode[K, V]) GetLeft() TraversableNode[K] {
	if bmn == nil || bmn.Left == nil {
		return nil
	}
	return bmn.Left
}

func (bmn *BinaryMapNode[K, V]) GetRight() TraversableNode[K] {
	if bmn == nil || bmn.Right == nil {
		return nil
	}
	return bmn.Right
}

func (bmn *BinaryMapNode[K, V]) GetValue() K {
	if bmn == nil {
		var zero K
		return zero
	}
	return bmn.Key
}

func (bmn *BinaryMapNode[K, V]) IsNil() bool {
	return bmn == nil
}
