package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Name:Lukman, Hobby:Play Game, Job:Software Engineer"
	res := parseString(str)

	for key, val := range res {
		fmt.Println(key, ":", val)
	}
}

func parseString(str string) map[string]string {
	// Initialize an empty map to 
	// store the key-value pairs.
	result := make(map[string]string)

	// Split the input string on 
	// the comma character to 
	// get a slice of key-value pairs.
	pairs := strings.Split(str, ",")
	
	// Iterate over the slice of pairs.
	for _, pair := range pairs {
		// Split the pair on the colon character 
		// to get the key and value.
		keyValue := strings.Split(pair, ":")
		// Check if the split resulted 
		// in a key and value.
		if len(keyValue) == 2 {
			// Trim leading and trailing white space 
			// from the key and value.
			key := strings.TrimSpace(keyValue[0])
			value := strings.TrimSpace(keyValue[1])
			// Add the key-value pair to the map.
			result[key] = value
		}
	}

	return result
}

