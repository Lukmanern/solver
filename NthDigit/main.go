package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Test Case 1
	n := 3
	fmt.Println("Test case 1:")
	fmt.Printf("n = %d\n", n)
	fmt.Printf("Output: %d\n", findNthDigit(n))
	fmt.Println("Expected Output : 3")

	// Test Case 2
	n = 11
	fmt.Println("Test case 2:")
	fmt.Printf("n = %d\n", n)
	fmt.Printf("Output: %d\n", findNthDigit(n))
	fmt.Println("Expected Output : 0")
	
	// Test Case 3
	n = 10000
	fmt.Println("Test case 3:")
	fmt.Printf("n = %d\n", n)
	fmt.Printf("Output: %d\n", findNthDigit(n))
	fmt.Println("Expected Output : 7")
}

func findNthDigit(n int) int {
	// Initialize the length of the current integer
	length := 1
	// Initialize the starting number of the current length
	start := 1

	// Iterate until n is within the range of the current length
	for n > length*9*start {
		// Subtract the number of digits in the current length from n
		n -= length * 9 * start

		// Increase the length of the integers
		length++
		// Update the starting number of the current length
		start *= 10
	}

	// Find the specific number in which the nth digit lies
	num := start + (n-1)/length

	// Convert the number to a string
	numStr := strconv.Itoa(num)

	// Find the nth digit and convert it to an integer
	return int(numStr[(n-1)%length] - '0')
}
