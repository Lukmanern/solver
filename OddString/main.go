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

func oddString(words []string) string {
	n := len(words[0])
	tmp := map[[19]int][]int{}
	for idx, word := range words {
		arr := [19]int{}
		for i := 0; i < n-1; i++ {
			arr[i] = int(word[i+1] - word[i])
		}
		tmp[arr] = append(tmp[arr], idx)
	}
	idx := -1
	for _, v := range tmp {
		if len(v) == 1 {
			idx = v[0]
			break
		}
	}
	return words[idx]
}