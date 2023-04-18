package main

import "fmt"

func main() {
	// arr = [1,3,5,7,9] -> take len(arr)-1 => 4 to sum
	// Our minimum sum is 1 + 3 + 5 + 7 = 16
	// and our maximum sum is 3 + 5 + 7 + 9 = 24
	// the output willbe 16 and 24

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(miniMaxSum(arr))
	// output 10 and 14

	arr = []int{1, 2, 3, 4, 5, 5}
	fmt.Println(miniMaxSum(arr))
	// output 15 and 19
}

func miniMaxSum(arr []int) (int, int) {
	// Return 0,0 if the array has less than 2 elements
	if len(arr) < 2 {
		return 0, 0
	}

	var min, max, total int
	// Find the minimum and maximum values in the array
	min, max = minMax(arr)

	// Calculate the total sum of the array
	for _, value := range arr {
		total += value
	}

	// Return the minimum sum by subtracting the maximum value from the total
	// and the maximum sum by subtracting the minimum value from the total
	return (total - max), (total - min)
}

// searching for bigest and
// lowest value form input
func minMax(arr []int) (int, int) {
	var max int = arr[0]
	var min int = arr[0]
	// loop through the array and update
	// the minimum and maximum values if necessary
	for _, value := range arr {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	return min, max
}
