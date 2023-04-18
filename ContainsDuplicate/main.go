package main

import "fmt"

func main() {
	dataset := [][]int{
		{1, 2, 3, 4, 5},  // false
		{2, 2, 3, 4, 5},  // true
		{3, 2, 3, 4, 5},  // true
		{11, 2, 3, 4, 5}, // false
		{-1, 2, 3, 4, 5}, // false
	}

	// loop through each sub-slice in the dataset
	for _, v := range dataset {
		fmt.Println(v, " : ", containsDuplicate(v))
	}
}

// check if the given slice
// contains any duplicates
func containsDuplicate(arr []int) bool {
	// i, j means index, v means value
	for i, v1 := range arr {
		for j, v2 := range arr {
			// skip comparing
			// the value with itself
			// or already comparing
			if i >= j {
				continue
			}
			if v1 == v2 {
				return true
			}
		}
	}

	return false
}
