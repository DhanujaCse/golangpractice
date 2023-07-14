package mapsassignment

import "sort"

func CountOccurences(arr []int) map[int]int {
	C := make(map[int]int)

	for _, num := range arr {
		C[num] = C[num] + 1 //updating the values by increasing value of occurences
	}
	return C
}
func RemoveDuplicates(arr []int) []int {
	C := make(map[int]int)
	for _, num := range arr {
		C[num] = 0
	}
	res := []int{}
	for k := range C {
		res = append(res, k) //appending the keys from C as keys are unique

	}

	return sort.IntSlice(res) //sort.IntSlice used to sort integer keys//.....sort.strings used to sort keys or values in string type

}
