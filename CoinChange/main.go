package main

import "fmt"

func main() {
	// Test case 1
	coins := []int{1, 2, 5}
	amount := 11
	result := coinChange(coins, amount)
	fmt.Println(result) // Expects 3

	// Test case 2
	coins  = []int{2}
	amount = 3
	result = coinChange(coins, amount)
	fmt.Println(result) // Expects -1

	// Test case 3
	coins  = []int{1}
	amount = 0
	result = coinChange(coins, amount)
	fmt.Println(result) // Expects 0

	// Test case 4
	coins  = []int{186, 419, 83, 408}
	amount = 6249
	result = coinChange(coins, amount)
	fmt.Println(result) // Expects 20

	// Test case 5
	coins  = []int{1, 3, 5}
	amount = 8
	result = coinChange(coins, amount)
	fmt.Println(result) // Expects 2

}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
