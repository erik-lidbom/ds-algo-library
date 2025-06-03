package linked_list

import (
	"fmt"
)

// INSERTION

func (linked_list *Linked_List) Prepend(value int)(*Single_Node, error){
	new_node := Single_Node{Value: value, Next: linked_list.Head}
	linked_list.Head = &new_node
	linked_list.Size++

	return &new_node, nil
}

func (linked_list *Linked_List) Append(value int) (*Single_Node, error) {
	new_node := Single_Node{Value: value, Next: nil}

	if linked_list.Head != nil{
		linked_list.Head = &new_node
		linked_list.Size++		
		return &new_node, nil
	}

	current_node := linked_list.Head

	for current_node != nil {
		current_node = current_node.Next

		if current_node.Next == nil {
			current_node.Next = &new_node
			break
		}
	}
	linked_list.Size++
	return &new_node, nil
}

func (linked_list *Linked_List) AddAt(index int, value int) (*Single_Node, error) {
	
	if index >= linked_list.Size{
		return nil, fmt.Errorf("index out of range")
	}


	new_node := Single_Node{Value: value, Next: nil}
	current_node := linked_list.Head
	var prev_node *Single_Node
	counter := 0

	for counter != index {
		prev_node = current_node
		current_node = current_node.Next
		counter++
	}

	if index == 0 {
		new_node.Next = linked_list.Head
		linked_list.Head = &new_node
	} else if counter == index {
		new_node.Next = current_node
		prev_node.Next = &new_node
	}

	linked_list.Size++

	return &new_node, nil
}
