package main

import "fmt"

type Dataset struct {
	nums    []int
	element int // removed element
}

func main() {
	/*
	res :
	[0 1 3 0 4 0 4 2] 5
	[1 2 3 4 6 8 7 7] 7
	*/
	dataset := []Dataset{
		{
			nums:    []int{0, 1, 2, 2, 3, 0, 4, 2},
			element: 2,
		},
		{
			nums:    []int{1, 2, 3, 4, 5, 6, 8, 7},
			element: 5,
		},
	}

	for _, d := range dataset {
		removeElement(d.nums, d.element)
	}
}

func removeElement(nums []int, val int) int {
	// 'c' represents the current 
	// index in the input slice
	c := 0

	// iterate over the input slice
	for _, num := range nums {
		// if the current number 
		// is not equal to the 
		// value to be removed
		if num != val {
			// replace the number at the 
			// current index with the 
			// current number
			nums[c] = num
			// increment the current index
			c++
		}
	}

	fmt.Println(nums, c)
	return c
}