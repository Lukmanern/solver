package main

import "math"

func divide(dividend int, divisor int) int {
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
