package linked_list
import ("fmt")



func (linked_list *Linked_List) PrintSize(){
	fmt.Printf("Current Size Of Linked List: %d\n", linked_list.Size)
}

func (linked_list *Linked_List) PrintNodes(){
	current_node := linked_list.Head
	counter := 0;
	for current_node != nil {
		fmt.Printf("Node at memory address %p, Value: %d, Index: %d\n", current_node, current_node.Value, counter)
		current_node = current_node.Next
		counter++
	}
}

func (linked_list *Linked_List) IsEmpty() bool {
	return linked_list.Size == 0 || linked_list.Head == nil
}

func (linked_list *Linked_List) GetSize() int {
	return linked_list.Size
}

func (linked_list *Linked_List) IndexOf(value int) int {
	if linked_list.Head == nil {
		return -1
	}
	current_node := linked_list.Head
	index_counter := 0

	for current_node != nil {
		if(current_node.Value == value) {
			return index_counter
		}
		current_node = current_node.Next
		index_counter++
	}

	return -1
}