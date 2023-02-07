package main

import "fmt"

func main() {
	input := [][]int{
		{1, 2, 3, 4, 5},
		{2, 3, 4, 5, 2},
	}
	output := flatten(input)
	fmt.Println(input, "->", output)  // Output: [1,2,3,4,5,2,3,4,5,2]
}

// flatten takes a 2D slice of integers 
// and returns a 1D slice of integers 
// that contains all elements 
// of the input slice.
func flatten(s [][]int) []int {
	// Create a 1D slice to store 
	// the flattened result.
	result := []int{}

	// Iterate through each sub-slice 
	// in the input 2D slice.
	for _, slice := range s {
		// Append all elements of the 
		// sub-slice to the result slice.
		result = append(result, slice...)
	}
	
	return result
}