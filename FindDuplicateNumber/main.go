package main

import "fmt"

func main() {
	dataset := [][]int{
		{1, 2, 3, 4, 1},
		{1, 2, 2, 3, 4},
		{3, 4, 5, 6, 6},
	}
	expected := []int{1, 2, 6}

	for i, d := range dataset {
		res := findDuplicate(d)
		fmt.Println(d, "->", res, "expect:", expected[i])
	}
}

func findDuplicate(nums []int) int {
	// Create a map to store the occurrences
	// of each element in the array
	hashMap := make(map[int]int, 0)
	// Initialize a variable to
	// store the duplicate value
	n := 0
	// Iterate through the
	// array of integers
	for _, n = range nums {
		// Increment the occurrence
		// count for the current element
		hashMap[n] += 1
		// Check if the current element
		// has occurred more than once
		if hashMap[n] >= 2 {
			// If so, break out of the loop
			// and return the current element
			break
		}
	}

	return n
}
