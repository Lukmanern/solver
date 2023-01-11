package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	i := 9000
	i = reverseInteger(i)
	fmt.Println(i) // 9
}

func reverseInteger(x int) int {
	if x == 0 {
		return 0
	}
	var isNeg bool = false
	var xString string = ""
	if x < 0 {
		isNeg = true
		x *= -1
	}

	for true {
		if x%100 == 0 {
			x /= 100
			continue
		}
		if x%10 == 0 {
			x /= 10
			continue
		}
		break
	}
	xString = strconv.Itoa(x)
	xString = reverseString(xString)
	x, _ = strconv.Atoi(xString)
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	if isNeg {
		x *= -1
	}
	return x
}

func reverseString(s string) string {
	// convert to rune
	rns := []rune(s)

	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
