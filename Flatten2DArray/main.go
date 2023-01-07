package main

import "fmt"

func main() {
	input := [][]int{
		{1, 2, 3, 4, 5},
		{2, 3, 4, 5, 2},
	}
	output := flatten(input)
	fmt.Println(input, "->", output)  // Output: [1,2,3,4,5,2,3,4,5,2]
}

func flatten(s [][]int) []int {
	result := []int{}
	for _, slice := range s {
		result = append(result, slice...)
	}
	return result
}
