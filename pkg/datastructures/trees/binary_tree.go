package trees

type BinaryNode struct {
	val any
	left *BinaryNode
	right *BinaryNode
}

type BinaryTree struct {
	root *BinaryNode
	size int
}