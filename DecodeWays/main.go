package main

import "fmt"

func main() {
	inputs := []string{"12", "226", "06","1","11","101"}
	expectedOutputs := []int{2, 3, 0, 1, 2, 1}

	for i, input := range inputs {
		fmt.Println(input, ":", numDecodings(input), "exp output :", expectedOutputs[i])
	}
}


func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	if s[0] == '0' {
		dp[1] = 0
	}
	for i := 2; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}
