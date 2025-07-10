package sort

import "fmt"

func BubbleSort (arr[]int) {
	size := len(arr)

	for i := 0; i < size - 1; i++ {
		for j := 0; j < size - i - 1; j++{
			if(arr[j] > arr[j + 1]){
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println("Newly sorted array")
	
	for i:=0; i < size - 1; i++{
		fmt.Printf("%d, ", arr[i])
	}
}
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

func InsertionSort (arr[]int) {

	size := len(arr)
	for i := 0; i < size - 1; i++ {
		j := i

		for j > 0 && arr[j] < arr[j - 1] {
			arr[j], arr[j - 1] = arr[j - 1], arr[j]
			j--
		}
	} 

	fmt.Println("Newly sorted array")
	
	for i := 0; i < size - 1; i++{
		fmt.Printf("%d, ", arr[i])
	}
}