package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(isAnagram("HAVANA", "AHNAVA"))
	fmt.Println(isAnagram("listen", "silent"))
	fmt.Println(isAnagram("elvis", "lives"))
	fmt.Println(isAnagram("debit card", "bad credit"))
	fmt.Println(isAnagram("schoolmaster", "the classroom"))
	fmt.Println(isAnagram("astronomer", "moon starer"))
	fmt.Println(isAnagram("funeral", "real fun"))
	fmt.Println(isAnagram("adoption", "no option")) // false 'd'
	fmt.Println(isAnagram("the eyes", "they see"))
}

func isAnagram(s string, t string) bool {
	// create two maps for each input string
	s_map := make(map[string]int)
	t_map := make(map[string]int)

	// iterate through each character
	// in first input string and add
	// it to the map
	for _, v := range s {
		// check if the character
		// is not empty or whitespace
		if string(v) == "" || string(v) == " " {
			continue
		}
		s_map[string(v)]++
	}

	// iterate through each character
	// in second input string
	// and add it to the map
	for _, v := range t {
		// check if the character
		// is not empty or whitespace
		if string(v) == "" || string(v) == " " {
			continue
		}
		t_map[string(v)]++
	}

	// print input strings to check
	fmt.Println(s + " & " + t + " :")

	// check if the maps are equal
	return reflect.DeepEqual(s_map, t_map)
}
