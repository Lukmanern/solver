package main

import (
	"fmt"
	"sort"
)

func main() {
	dataset := [][]int{
		{4, 3, 1, 3, 6,}, // 2
		{2, 3, 4, 6, 7, 2, 10}, // 1
		{1, 2, 3, 4, 5, 6, 8, 7}, // 9
	}

	for _, d := range dataset {
		fmt.Println(d, ":", firstMissingPositive(d))
	}
}

// firstMissingPositive returns the smallest positive 
// that is missing from the input slice 'nums'
func firstMissingPositive(nums []int) int {
	// sort the input slice in ascending order
	sort.Ints(nums)

	// initialize the smallest missing 
	// positive integer with 1
	c := 1

	// iterate over the sorted slice
	for _, num := range nums {
		// if the current number is greater 
		// than 'c', we have found the 
		// first missing positive integer
		if num > c {
			break
		}
		// if the current number is equal 
		// to 'c', increment 'c' by 1
		if num == c {
			c++
		}
	}

	return c
}
