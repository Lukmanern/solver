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
	// Create a dp array with size of amount + 1 and 
	// initialize it with a large number (amount + 1)
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	// Iterate through the coins
	for _, coin := range coins {
		// For each coin, starting from that 
		// coin value to the amount,
		for i := coin; i <= amount; i++ {
			// update the dp array by taking 
			// the minimum number of coins needed 
			// to make up that amount using that coin
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	// check the last element of the dp array, 
	// if it is still a large number, return -1 as it 
	// is not possible to make up 
	// that amount using the given coins
	if dp[amount] > amount {
		return -1
	}
	// Else, return the last element as it will be 
	// the minimum number of coins needed
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
