package linked_list

import (
	"fmt"
)

// DELETION

func (linked_list *Linked_List) RemoveFirst() (int, error) {
	if linked_list.Head != nil {
		return 0, fmt.Errorf("list is empty")
	}
	
	removed_head := linked_list.Head
	new_head := removed_head.Next
	linked_list.Head = new_head
	linked_list.Size--

	return removed_head.Value, nil
}

func (linked_list *Linked_List) RemoveLast()(int, error){
	if linked_list.Head != nil {
		fmt.Println("Linked List is empty")
		return 0, fmt.Errorf("list is empty")
	}

	current_node := linked_list.Head
	var prev_node *Single_Node

	for current_node.Next != nil {
		prev_node = current_node	
		current_node = current_node.Next
	}

	prev_node.Next = nil
	linked_list.Size--

	return current_node.Value, nil
}

func (linked_list *Linked_List) RemoveAt(index int) (int, error){
	
	if(index < 0 || index >= linked_list.Size){
		return 0, fmt.Errorf("index out of range")
	}

	if(index == 0){
		removed_val := linked_list.Head.Value
		linked_list.Head = linked_list.Head.Next
		linked_list.Size--
		return removed_val, nil
	}

	current_node := linked_list.Head

	for i := 0; i < index - 1; i++ {
		current_node = current_node.Next
	}

	removed_node := current_node.Next
	current_node.Next = removed_node.Next
	linked_list.Size--

	return removed_node.Value, nil
}

func (linked_list *Linked_List) Clear() (error) {

	if(linked_list.Size == 0){
		return fmt.Errorf("list already cleared")
	}

	linked_list.Head = nil
	linked_list.Size = 0
	return nil
}