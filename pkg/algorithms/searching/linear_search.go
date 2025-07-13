package search

/*
Algorithm to search for elements in an unordered list
Complexity: O(n)
*/
func LinearSearch (arr[] int, element int) (int,bool) {
	
	for i := range len(arr){
		if(arr[i] == element){
			return i, true
		}
	}
	return 0, false
}