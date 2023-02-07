package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	merge(nums1, m, nums2, n)
	fmt.Println(nums1) // result : 1 2 2 3 5 6
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	// i: pointer for nums1, j: pointer 
	// for nums2, k: pointer for merged array
	i, j, k := m-1, n-1, m+n-1
	// compare and merge elements 
	// from nums1 and nums2
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	// add remaining elements from 
	// nums2 to merged array
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
