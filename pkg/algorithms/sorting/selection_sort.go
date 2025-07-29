package sorting

import (
	"cmp"

	"ds-algorithms/pkg/datastructures/common"
)

func SelectionSort[T cmp.Ordered](arr common.Searchable[T]) {
	size := arr.Size()

	for i := 0; i < size-1; i++ {
		minIndex := i

		for j := i + 1; j < size; j++ {
			currVal, _ := arr.Get(j)
			minVal, _ := arr.Get(minIndex)
			if currVal < minVal {
				minIndex = j
			}
		}
		currVal, _ := arr.Get(i)
		minVal, _ := arr.Get(minIndex)
		arr.Set(i, minVal)
		arr.Set(minIndex, currVal)
	}
}
