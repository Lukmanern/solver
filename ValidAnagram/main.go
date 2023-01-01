package main

import (
	"fmt"
	"reflect"
)

func isAnagram(s string, t string) bool {
	s_map := make(map[string]int)
	t_map := make(map[string]int)

	for _, v := range s {
		if string(v) == "" || string(v) == " " {
			continue
		}
		s_map[string(v)]++
	}

	for _, v := range t {
		if string(v) == "" || string(v) == " " {
			continue
		}
		t_map[string(v)]++
	}

	fmt.Println(s + " & " + t + " :")

	return reflect.DeepEqual(s_map, t_map)
}

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