package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/valid-parentheses/
func main(){
	// dataSet is a slice of strings 
	// that contains test cases
	dataSet := []string{
		"[]{()[][]}", // true
		"[]{()[(][]}", // false
		"[]{[]()[][]}", // true
		"[](){()[][]}", // true
		"[]{()[{][]}", // false
	}

	// range over the dataSet and 
	// pass each test case to isValid
	// print the result of isValid for each test case
	for _, v := range dataSet {
		fmt.Println(isValid(v))
	}
}

// isValid checks if the given string is valid
// a string is considered valid if all 
// its brackets are properly closed
func isValid(s string) bool {
	// len is the length of the given string
	len := len(s)
	// if the length of the string is odd, 
	// it cannot be valid, return false
	if len%2 != 0 {return false}
	
	// loop through the string 
	// until all brackets are closed
	for i := 0; i < len; i++ {
		// replace all instances of () or 
		// {} or [] with an empty string
		s = strings.Replace(s, "()", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		s = strings.Replace(s, "{}", "", -1)
	}
	
	// if the string is empty, 
	// all brackets have been closed
	// return true, else return false
	return s == ""
}
