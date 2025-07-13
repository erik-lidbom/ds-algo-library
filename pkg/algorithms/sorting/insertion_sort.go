package sorting

import ("fmt")

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