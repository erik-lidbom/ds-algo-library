package array

import (
	"errors"
	"fmt"
	"strings"
)

type ArrayList struct {
	arr []any
	size int
}

func NewArrayList() *ArrayList {
	return &ArrayList{arr: make([]any, 10), size: 0}
}

func (al *ArrayList) Size() int {
	return al.size
}

func (al *ArrayList) IsEmpty() bool {
	return al.size == 0
}

func (al *ArrayList) Add(index int, elem any) (error) {
	if index < 0 || index > al.size {
		return  errors.New("index out of bounds")
	}

	if al.size >= len(al.arr){
		al.resize()
	}

	if index == al.size {
		al.arr[index] = elem
		al.size++
		return nil
	}
	
	for i := al.size; i > index; i-- {
		al.arr[i] = al.arr[i - 1]
	}
	al.arr[index] = elem
	al.size++
	
	return nil
}
func (al *ArrayList) Get(index int) (any, error) {
	if index < 0 || index >= al.size {
		return nil, errors.New("index out of bounds")
	}

	return al.arr[index], nil
}
func (al *ArrayList) Set(index int, elem any) (error) {
	if index < 0 || index >= al.size {
		return errors.New("index out of bounds")
	}

	al.arr[index] = elem
	return nil
}
func (al *ArrayList) Remove(index int) (any, error) {

	if index < 0 || index >= al.size {
		return nil, errors.New("index out of bounds")
	}

	removedVal := al.arr[index]

	for i := index + 1; i < al.size; i++ {
		al.arr[i - 1] = al.arr[i]
	}
	al.arr[al.size - 1] = nil
	al.size--

	if al.size * 3 <= len(al.arr) {
		al.shrink()
	}
	return removedVal, nil
}

func (al *ArrayList) String() string {
    if al.size == 0 {
        return "ArrayList: [] (Size: 0)"
    }

    var sb strings.Builder
    sb.WriteString("ArrayList: [")
    for i := 0; i < al.size; i++ {
        sb.WriteString(fmt.Sprintf("%v", al.arr[i]))
        if i < al.size-1 {
            sb.WriteString(", ")
        }
    }
    sb.WriteString(fmt.Sprintf("] (Size: %d)", al.size))
    return sb.String()
}

func (al *ArrayList) resize() {
	oldArray := al.arr

	al.arr = make([]any, len(al.arr) * 2)
	for i := 0; i < al.size; i++ {
		al.arr[i] = oldArray[i]
	}
}
func (al *ArrayList) shrink() {
	oldArray := al.arr
	al.arr = make([]any, len(al.arr) / 2)

	for i := 0; i < al.size; i++ {
		al.arr[i] = oldArray[i]
	}
}