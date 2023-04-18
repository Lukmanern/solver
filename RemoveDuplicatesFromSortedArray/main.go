package main

import "fmt"

func main() {
	dataset := [][]int{
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, // 5 : 0, 1, 2, 3, 4, _, _, _, _, _,
	}

	for _, d := range dataset {
		removeDuplicates(d)
	}
}

func removeDuplicates(nums []int) int {
	// 'c' represents the current
	// index in the input slice
	c := 0

	// iterate over the input slice
	for _, num := range nums {
		// if the current number is
		// different from the number
		// at the current index
		if num != nums[c] {
			// increment the current index
			c++
			// replace the number at
			// the current index
			// with the current number
			nums[c] = num
		}
	}

	fmt.Println(c, ":", nums)
	return c + 1
}
