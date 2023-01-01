package main

import "fmt"

func main() {
	str := "1234567899"
	fmt.Println(reverseString(str)) // 9987654321
}

func reverseString(s string) string {
	// If the length of the string is 0, 
	// return the string as is.
	if len(s) == 0 {
		return s
	}
	
	// Add the first character of s 
	// to the end of the reversed substring.
	// This process continues until 
	// the entire string has been reversed.
	return reverseString(s[1:]) + s[0:1]
}