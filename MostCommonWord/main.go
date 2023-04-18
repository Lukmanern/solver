package main

import (
	"fmt"
	"strings"
)

type Dataset struct {
	paragraph string
	banned    []string
}

func main() {
	// result : sky and wind
	dataset := []Dataset{
		{
			paragraph: "The sky is blue, the sky is clear. The sun is shining and the sky is a beautiful shade of blue. The clouds are white and the sky is blue.",
			banned:    []string{"the", "and", "is"},
		},
		{
			paragraph: "The wind is blowing, the trees are swaying. The leaves rustle as the wind blows, creating a symphony of sound. The wind is strong, but the trees stand tall. Despite the wind, the trees stand firm and continue to sway with the wind's rhythm.",
			banned:    []string{"the", "leaves", "rustle"},
		},
	}

	for _, data := range dataset {
		fmt.Println(mostCommonWord(data.paragraph, data.banned))
	}
}

// mostCommonWord returns the most common word
// in a given paragraph, excluding
// the words in the banned list.
func mostCommonWord(paragraph string, banned []string) string {
	// Convert the paragraph to lowercase.
	s := strings.ToLower(paragraph) + " "

	// Create a map to store the
	// frequency of each word.
	str := make(map[string]int)

	// s1 is a variable to store the
	// current word as it is being processed.
	s1 := ""

	// Iterate through each character
	// in the lowercase paragraph.
	for _, v := range s {
		// If the character is a space or a punctuation mark,
		// it marks the end of a word, so add it to the frequency map.
		if v == ' ' || strings.Contains("!?',;.", string(v)) {
			if len(s1) != 0 {
				str[s1]++
				s1 = ""
			}
		} else {
			// If the character is not a space
			// or punctuation mark, add it to s1.
			s1 += string(v)
		}
	}

	// Remove the banned words
	// from the frequency map.
	for _, v := range banned {
		delete(str, v)
	}

	// res is a variable to store
	// the most common word.
	res := ""

	// max is a variable to store
	// the maximum frequency.
	max := 0

	// Iterate through the frequency map
	// to find the word with
	// the highest frequency.
	for k, v := range str {
		if v > max {
			max = v
			res = k
		}
	}

	// Return the most common word.
	return res
}
