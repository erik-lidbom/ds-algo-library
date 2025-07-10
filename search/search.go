package search

/*
Algorithm to search for elements in an unordered list
Complexity: O(n)
*/
func SequentialSearch (arr[] int, element int) (int,bool) {
	
	for i := range len(arr){
		if(arr[i] == element){
			return i, true
		}
	}
	return 0, false
}


/*
Searching algorithm of an ordered list
Complexity: O(log n)
*/

func BinarySearch(arr []int, element int) (int, bool){

	left_index := 0
	right_index := len(arr) - 1

	for left_index <= right_index {
		mid_index := left_index + (right_index - left_index) / 2

		if(element == arr[mid_index]){
			return mid_index, true
		} else if(element < arr[mid_index]){
			right_index = mid_index - 1
		} else if(element > arr[mid_index]){
			left_index = mid_index + 1
		}
	}
	return 0, false
}