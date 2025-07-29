package sorting

import (
	"cmp"

	"ds-algorithms/pkg/datastructures/common"
)

func BubbleSort[T cmp.Ordered](arr common.Searchable[T]) {
	size := arr.Size()

	for i := 0; i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
			valX, _ := arr.Get(j)
			valY, _ := arr.Get(j + 1)

			if valX > valY {
				arr.Set(j, valY)
				arr.Set(j+1, valX)
			}
		}
	}
}
