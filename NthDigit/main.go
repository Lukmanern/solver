package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findNthDigit(1000)) // expect 3
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
