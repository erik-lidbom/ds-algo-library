package heap

func isLeaf(pos, size int) bool {
	return pos >= size/2
}

func getLeftChildIndex(pos int) int {
	return 2*pos + 1
}

func getRightChildIndex(pos int) int {
	return 2*pos + 2
}

func getParent(pos int) int {
	return int((pos - 1) / 2)
}
