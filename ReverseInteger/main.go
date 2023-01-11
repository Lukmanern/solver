package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	i := 9000
	i = reverseInteger(i)
	fmt.Println(i) // 9

	i = -809011
	i = reverseInteger(i)
	fmt.Println(i) // 110908
}

func reverseInteger(x int) int {
    	// Check if the input is 0
	if x == 0 {
		return 0
	}
    	// Initialize a variable to check 
	// if the input is negative
	var isNeg bool = false
	var xString string = ""
    	// If the input is negative, set isNeg 
	// to true and make x positive
	if x < 0 {
		isNeg = true
		x *= -1
	}

	// Remove trailing 0s from x
	// For example, if x = 200, it will become 2
	for true {
		if x%100 == 0 {
			x /= 100
			continue
		}
		if x%10 == 0 {
			x /= 10
			continue
		}
		break
	}
    	// Convert x to a string
	xString = strconv.Itoa(x)
    	// Reverse the string
	xString = reverseString(xString)
    	// Convert the reversed 
	// string back to an int
	x, _ = strconv.Atoi(xString)
    	// Check if x is within the 
	// 32-bit signed integer range
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
    	// If the original input was negative, 
	// make x negative again
	if isNeg {
		x *= -1
	}
	return x
}



func reverseString(s string) string {
	// convert to rune
	rns := []rune(s)

	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
