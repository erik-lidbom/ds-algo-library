package sorting

import (
	"cmp"
	"ds-algorithms/pkg/datastructures/common"
)

func InsertionSort[T cmp.Ordered] (
	arr common.Searchable[T]) {

	size := arr.Size()

	for i := 1; i < size; i++ {
		j := i
		
		valX, _ := arr.Get(j)

		for j > 0 {
			valY, _ := arr.Get(j - 1)
			if valX >= valY {
				break
			}
			arr.Set(j, valY)
			j--
		}
		arr.Set(j, valX)
	} 
}