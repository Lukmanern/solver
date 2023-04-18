package main

import "fmt"

func main() {
	dataset := []int{
		1, 2, 0, 2, 2, 1, 0, 1, 2, 0,
	}
	sortColors(dataset)

	fmt.Println(dataset)
}

func sortColors(nums []int) {
	// Initialize pointers for left
	// and right sides of the array
	left, right := 0, len(nums)-1

	// Initialize pointer
	// for current element
	idx := 0

	// Loop through the array until
	// current pointer reaches the end
	for idx <= right {
		if nums[idx] == 0 {
			// Swap current element with element
			// at left pointer if it's 0
			nums[idx], nums[left] = nums[left], nums[idx]

			// Move left pointer to the right
			left++

			// Move current pointer to the right
			idx++
		} else if nums[idx] == 2 {
			// Swap current element with element
			// at right pointer if it's 2
			nums[idx], nums[right] = nums[right], nums[idx]

			// Move right pointer to the left
			right--
		} else {
			// Move current pointer to the
			// right if element is 1
			idx++
		}
	}
}
