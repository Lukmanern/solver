package main

import "fmt"

func main() {
	arr := []int{10, 5, 3, 2, 6, 8, 1, 7, 9, 11, 13} // 4 and 12
	missingNumbers := findMissingNumbers(arr)
	fmt.Printf("Multiple numbers\nThe missing numbers is: %d\n", missingNumbers)

	arr = []int{10, 5, 3, 2, 6, 8, 1, 7, 9, 11, 13, 4} // 12
	missingNumber := findMissingNumber(arr)
	fmt.Printf("Single number\nThe missing numbers is: %d\n", missingNumber)
}

func findMissingNumber(arr []int) int {
	// Find the sum of all the numbers in the array
	var sum int
	for _, value := range arr {
		sum += value
	}

	// Calculate the sum of all the numbers 
	// from 1 to the length of the array + 1. 
	// This is the expected sum if there 
	// are no missing numbers
	expectedSum := (len(arr) + 1) * (len(arr) + 2) / 2
	
	// Return the difference between 
	// the expected sum and the actual
	// sum as the missing number
	return expectedSum - sum
}

func findMissingNumbers(arr []int) []int {
	// Create a map to store 
	// the counts of each 
	// number in the array
	counts := make(map[int]int)
	for _, value := range arr {
		counts[value]++
	}

	// Create a slice to store 
	// the missing numbers
	var missingNumbers []int

	// Find the missing numbers by iterating 
	// through all the numbers from 1 to the 
	// length of the array + 1 and checking 
	// if they are in the counts map
	for i := 1; i <= len(arr)+1; i++ {
		if _, ok := counts[i]; !ok {
			missingNumbers = append(missingNumbers, i)
		}
	}

	return missingNumbers
}
