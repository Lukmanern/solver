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
	c := 0

	for _, num := range nums {
		if num != val {
			nums[c] = num
			c++
		}
	}

	fmt.Println(nums, c)
	return c
}