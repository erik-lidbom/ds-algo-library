package linked_list

import (
	"fmt"
)

// Access & Retrieval
func (linked_list *Linked_List) PeekFirst() (int, error) {
	if(linked_list.Head != nil){
		return 0, fmt.Errorf("list is empty")
	}

	return linked_list.Head.Value, nil
}

func (linked_list *Linked_List) PeekLast() (int, error){
	if(linked_list.Head != nil){
		return 0, fmt.Errorf("list is empty")
	}

	current_node := linked_list.Head

	for current_node.Next != nil {
		current_node = current_node.Next
	}

	return current_node.Value, nil
}

func (linked_list *Linked_List) Get(index int) (int, error) {
	if linked_list.Head != nil {
		return 0, fmt.Errorf("list is empty")
	}

	if(index < 0 || index >= linked_list.Size){
		return 0, fmt.Errorf("index out of range")
	}

	if index == 0 {
		return linked_list.Head.Value, nil
	}

	current_node := linked_list.Head
	
	for i := 0; i < index ; i++ {
		current_node = current_node.Next
	}

	return current_node.Value, nil
}