package main

import (
	"fmt"
	"sort"
)

func main() {
	dataset := [][]int{
		{4, 3, 1, 3, 6,}, // 2
		{2, 3, 4, 6, 7, 2, 10}, // 1
		{1, 2, 3, 4, 5, 6, 8, 7}, // 9
	}

	for _, d := range dataset {
		fmt.Println(d, ":", firstMissingPositive(d))
	}
}

func firstMissingPositive(nums []int) int {
    sort.Ints(nums)
    c := 1
    for _, num := range nums {
        if num > c {
            break
        }
        if num == c {
            c++
        }
    }

    return c
}