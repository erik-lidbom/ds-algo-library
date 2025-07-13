package sorting

import "fmt"

func SelectionSort (arr[]int) {
	size := len(arr)

	for i := 0; i < size - 1; i++ {
		min_index := i
		
		for j := i + 1; j < size; j++ {
			if(arr[j] < arr[min_index]){
				min_index = j
			}
		}

		arr[i], arr[min_index] = arr[min_index], arr[i]
	}

	fmt.Println("Newly sorted array")
	
	for i:=0; i < size - 1; i++{
		fmt.Printf("%d, ", arr[i])
	}
}