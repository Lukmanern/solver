package main

import "fmt"

func main() {
	dataset := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4}, // 6 = 4, -1, 2, 1
		{1, 2, 3, -4, 5, -6, 7, 8}, // 16 = all-sum
		{-1, -2, -3, -4, -5, -6, -7, -8}, // -1 = -1
		{5, 4, -1, 7, -8, -3, 11, -1, -6}, // 15 = 5, 4, -1, 7, -8, -3, 11
		{-4, 2, -5, 1, 2, 3, 6, -5, 2}, // 10 = 2, 1, 2, 3, 6
	}

	for _, v := range dataset {
		fmt.Println(v, ":", maxSubArray(v))
	}
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum = findMax(nums[i], currentSum+nums[i])
		maxSum = findMax(maxSum, currentSum)
	}
	return maxSum
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}


