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
	maxSum := 0
	maxIndex := 0
	for i, row := range arr {
		rowSum := 0
		for _, value := range row {
			rowSum += value
		}
		if rowSum > maxSum {
			maxSum = rowSum
			maxIndex = i
		}
	}

	return maxIndex, maxSum
}
