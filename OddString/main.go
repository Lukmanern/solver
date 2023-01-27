package main

import "fmt"

func main() {
	dataset := [][]string{
		{"adc", "wzy", "abc"},        // abc
		{"aaa", "bob", "ccc", "ddd"}, // bob
	}

	for i := range dataset {
		fmt.Println(oddString(dataset[i]))
	}
}

// oddString finds the string in the input 
// slice that has unique character difference
// It returns the unique string
func oddString(words []string) string {
	// n represents the length of 
	// the first word in the input slice
	n := len(words[0])
	// tmp is a map that stores the 
	// difference of each character 
	// of a word as a key and 
	// the index of that word as value
	tmp := map[[19]int][]int{}
	for idx, word := range words {
		// arr stores the difference 
		// of each character of the current word
		arr := [19]int{}
		for i := 0; i < n-1; i++ {
			arr[i] = int(word[i+1] - word[i])
		}
		// appends the index of the 
		// current word to the value of the key
		tmp[arr] = append(tmp[arr], idx)
	}
	// idx stores the index 
	// of the unique word
	idx := -1
	for _, v := range tmp {
		if len(v) == 1 {
			idx = v[0]
			break
		}
	}
	return words[idx]
}