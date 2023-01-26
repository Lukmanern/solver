package main

import (
	"fmt"
	"strings"
)

func main() {
	s := " a stone    in a mountain "
	// short way
	fmt.Println(len(strings.Fields(s)))
	// gentlemen way
	fmt.Println(wordsCounter(s))
}

func wordsCounter(s string) int {
	// Initialize the word count to 0.
	count := 0
	// Initialize a flag to keep track of whether 
	// or not we are currently in a word.
	inWord := false
	// Iterate over the characters in the string.
	for _, c := range s {
		if c != ' ' {
			// If we are not 
			// currently in a word
			if !inWord {
				// Increment the word count and 
				// set the inWord flag to true.
				count++
				inWord = true
			}
		} else {
			// Set the inWord flag to false.
			inWord = false
		}
	}

	return count
}
