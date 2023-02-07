package main

import (
	"fmt"
	"sort"
)

func main() {
	dataset := [][]int{
		{1,2,3,4,5},
		{3,4,5},
	}
	fmt.Println(findMedianSortedArrays(dataset[0], dataset[1]))
}

// findMedianSortedArrays is a function that takes 
// in two sorted slices of integers and 
// returns the median value of the merged slice
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// merge the two slices
	var merged []int = append(nums1, nums2...)
	// get the length of the merged slice
	var len int = len(merged)

	// sort the merged slice
	sort.Slice(merged, func(i, j int) bool {
		return merged[i] < merged[j]
	})

	// if the length of the merged slice is odd, 
	// return the value at the middle index
	if len%2 != 0 {
		return float64(merged[len/2])
	}

	// if the length of the merged slice is even, return 
	// the average of the values at the middle indices
	var midValue int = merged[(len/2)-1] + merged[len/2]
	
	return float64(midValue) / 2
}
