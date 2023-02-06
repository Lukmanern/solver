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
	// result sky and wind
	dataset := []Dataset{
		{
			paragraph: "The sky is blue, the sky is clear. The sun is shining and the sky is a beautiful shade of blue. The clouds are white and the sky is blue.",
			banned: []string{"the", "and", "is"},
		},
		{
			paragraph: "The wind is blowing, the trees are swaying. The leaves rustle as the wind blows, creating a symphony of sound. The wind is strong, but the trees stand tall. Despite the wind, the trees stand firm and continue to sway with the wind's rhythm.",
			banned: []string{"the", "leaves", "rustle"},
		},
	}

	for _, data := range dataset {
		fmt.Println(mostCommonWord(data.paragraph, data.banned))
	}
}

func mostCommonWord(paragraph string, banned []string) string {
	s := strings.ToLower(paragraph) + " "

	str := make(map[string]int)
	s1 := ""
	for _, v := range s {
		if v == ' ' || strings.Contains("!?',;.", string(v)) {
			if len(s1) != 0 {
				str[s1]++
				s1 = ""
			}
		} else {
			s1 += string(v)
		}
	}

	for _, v := range banned {
		delete(str, v)
	}

	res := ""
	max := 0
	for k, v := range str {
		if v > max {
			max = v
			res = k
		}
	}

	return res
}