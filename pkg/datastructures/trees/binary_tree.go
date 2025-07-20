package trees

import "fmt"

type BinaryNode struct {
	val any
	left *BinaryNode
	right *BinaryNode
}

type BinaryTree struct {
	root *BinaryNode
	size int
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{root: nil, size: 0}
}

func (bt *BinaryTree) Size() int {
	return bt.size
}

func (bt *BinaryTree) IsEmpty() bool {
	return bt.size == 0
}

func (bt *BinaryTree) PreOrder(node *BinaryNode) {
	if node != nil {
		fmt.Printf("Node: %v\n", node)
		bt.PreOrder(node.left)
		bt.PreOrder(node.right)
	}
}

func (bt *BinaryTree) PostOrder(node *BinaryNode) {
	if node != nil {
		bt.PostOrder(node.left)
		bt.PostOrder(node.right)
		fmt.Printf("Node %v\n", node)
	}
}

func (bt *BinaryTree) InOrder(node *BinaryNode) {
	if node != nil {
		bt.InOrder(node.left)
		fmt.Printf("Node: %v\n", node)
		bt.InOrder(node.right)
	}
}

func (bt *BinaryTree) Clear() error {
	bt.root = nil
	bt.size = 0
	return nil
}

func (bt *BinaryTree) IsLeaf(node *BinaryNode) bool {
	return node.left == nil && node.right == nil
}