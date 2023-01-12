package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// Initialize an empty slice 
	// to store the result &
	// sort the input slice
	var res [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		// Skip the current element 
		// if it's a duplicate
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// Init left and right pointers
		l, r := i+1, len(nums)-1
		// Iterate through the remaining elements 
		// to the right of the current element
		for l < r {
			// Calculate the sum of 
			// the current triplet
			s := nums[i] + nums[l] + nums[r]
			// If the sum is less than 0, 
			// increment the left pointer
			if s < 0 {
				l++
			} else if s > 0 {
				// If the sum is greater than 0, 
				// decrement the right pointer
				r--
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4, 5}
	fmt.Println(threeSum(nums))
}
