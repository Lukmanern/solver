package main

import (
	"fmt"
	"strconv"
)

func createPalindrome(n int) string {
	// Create an empty string
	// for the palindrome
	p := ""
	// Append the numbers from 1
	// to n to the palindrome string
	for i := 1; i <= n; i++ {
		p += strconv.Itoa(i)
	}
	// Append the numbers from n-1
	// down to 1 to the palindrome string
	for i := n - 1; i > 0; i-- {
		p += strconv.Itoa(i)
	}
	// Return the palindrome
	return p
}

func main() {
	fmt.Println(createPalindrome(5)) // Output: "123454321"
}
