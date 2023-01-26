package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4} // [1, 3, 6, 10]
	result := runningSum(nums)
	fmt.Println(result)
}

// runningSum is a function that 
// takes in a slice of integers
// and returns a new slice with 
// the running sum of the original slice.
func runningSum(nums []int) []int {
	// Iterate through the 
	// input slice starting at index 1
	for i := 1; i < len(nums); i++ {
		// Add the current value to the previous 
		// value and update the current index
		nums[i] += nums[i-1]
	}
	return nums
}
