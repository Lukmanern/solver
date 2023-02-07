package main

import "fmt"

func main() {
	arr := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		{1, 2, 3, -4, 5, -6, 7, 8}, // 16 
		{-1, -2, -3, -4, -5, -6, -7, -8},
		{5, 4, -1, 7, -8, -3, 11, -1, -6},
		{-4, 2, -5, 1, 2, 3, 6, -5, 2},
	}
	maxIndex, maxSum := maxSub2DArray(arr)
	fmt.Println("Index:", maxIndex, "Sum:", maxSum)
}

func maxSub2DArray(arr [][]int) (int, int) {
	// Initialize maxSum and maxIndex 
	// to keep track of the maximum 
	// sum and its corresponding index.
	maxSum := 0
	maxIndex := 0

	// Iterate through each element in the 
	// arr, which is a slice of integers.
	for i, row := range arr {
		// Initialize rowSum to keep track 
		// of the sum of elements in the current row.
		rowSum := 0

		// Iterate through each element in the 
		// current row and add it to rowSum.
		for _, value := range row {
			rowSum += value
		}

		// If it is greater, 
		// update maxSum and maxIndex.
		if rowSum > maxSum {
			maxSum = rowSum
			maxIndex = i
		}
	}

	// return the maxIndex and maxSum
	return maxIndex, maxSum
}

