package main

import (
	"fmt"
	"math"
)

func main() {
	var res_1, res_2 int
	dataset := [][2]int{
		{10, 3},
		{11, 3},
		{-1, 3},
		{-10, 3},
		{0, 3},
	}

	for _, data := range dataset {
		res_1 = divide_1(data[0], data[1])
		res_2 = divide_2(data[0], data[1])
		fmt.Println("divide_1 :", res_1)
		fmt.Println("divide_2 :", res_2)
		fmt.Println() // enter

		if res_1 != res_2 {
			panic("error : result of divide_1 != divide_2")
		}
	}
}

func divide_1(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	var result int = dividend / divisor

	max32 := math.MaxInt32
	min32 := math.MinInt32

	if result >= max32 {
		return max32
	} else if result <= min32 {
		return min32
	}

	return result
}

func divide_2(dividend int, divisor int) int {
	// Handle division by zero
	if divisor == 0 {
		return math.MaxInt32
	}
	if dividend == 0 {
		return 0
	}

	// Handle negative numbers
	isNegative := (dividend < 0) != (divisor < 0)
	absDividend := int(math.Abs(float64(dividend)))
	absDivisor := int(math.Abs(float64(divisor)))

	// Initialize variables for the result and the current power of two
	result := 0
	powerOfTwo := 1

	// Keep doubling the divisor until it's greater than the dividend
	for absDivisor <= absDividend {
		absDivisor <<= 1
		powerOfTwo <<= 1
	}

	// Subtract the doubled divisor from the dividend and update the result
	for powerOfTwo > 1 {
		absDivisor >>= 1
		powerOfTwo >>= 1
		if absDividend >= absDivisor {
			absDividend -= absDivisor
			result += powerOfTwo
		}
	}

	// Handle overflow and return the result
	if isNegative {
		result = -result
	}
	if result > math.MaxInt32 {
		return math.MaxInt32
	}
	if result < math.MinInt32 {
		return math.MinInt32
	}
	return result
}
