package main

import "fmt"

func main() {
	inputs := []string{"12", "226", "06", "1", "11", "101"}
	expectedOutputs := []int{2, 3, 0, 1, 2, 1}

	for i, input := range inputs {
		fmt.Println(input, ":", numDecodings(input), "exp output :", expectedOutputs[i])
	}
}

func numDecodings(s string) int {
	// get the length of the string
	n := len(s)
	// if the string is empty, return 0
	if n == 0 {
		return 0
	}
	// create a dp array with n+1 elements
	// and initialize the first two elements
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	// if the first element of the
	// string is '0', set dp[1] to 0
	if s[0] == '0' {
		dp[1] = 0
	}
	// loop through the string starting
	// from the 3rd element
	for i := 2; i <= n; i++ {
		// if the current element is not '0', add
		// the number of ways to decode the substring
		// without the current element
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		// if the current element and the previous element
		// form a valid combination ('1' or '2' followed by
		// a number <= '6'), add the number of ways to decode
		// the substring without the current and previous elements
		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6') {
			dp[i] += dp[i-2]
		}
	}
	// return the last element of
	// the dp array which is
	// the number of ways
	// to decode the entire string
	return dp[n]
}
