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
	result := make(map[string]string)

	pairs := strings.Split(str, ",")
	for _, pair := range pairs {
		keyValue := strings.Split(pair, ":")
		if len(keyValue) == 2 {
			// trims space " "
			result[strings.TrimSpace(keyValue[0])] = strings.TrimSpace(keyValue[1])
		}
	}

	return result
}
