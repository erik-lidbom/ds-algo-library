package linked_list

type Single_Node struct  {
	Value int
	Next *Single_Node
}

type Linked_List struct {
	Head *Single_Node
	Size int
}