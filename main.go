package main

import (
	"os"
	"data-structures/linked_list")

type Single_Node struct  {
	value int
	next *Single_Node
}

type Linked_List struct {
	head *Single_Node
	size int
}

func main(){

    myList := &linked_list.Linked_List{} 
    myList.Prepend(10)
    myList.Append(20)
	
	os.Exit(0)
}

