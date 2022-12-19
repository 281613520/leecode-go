package main

func countVowelPermutation(n int) int {
	const mod int = 1e9 + 7
	dp := make([][]int, n)

	for i := 0; i < 5; i++ {
		dp[i] = make([]int, 5)
	}

	dp[0][0] = 1
	dp[0][1] = 1
	dp[0][2] = 1
	dp[0][3] = 1
	dp[0][4] = 1

	for i := 1; i < n; i++ {
		dp[i][0] = (dp[i-1][1] + dp[i-1][2] + dp[i-1][4]) % mod
		dp[i][1] = (dp[i-1][0] + dp[i-1][2]) % mod
		dp[i][2] = (dp[i-1][1] + dp[i-1][3]) % mod
		dp[i][3] = (dp[i-1][2]) % mod
		dp[i][4] = (dp[i-1][2] + dp[i-1][3]) % mod
	}

	ans := 0

	for i := 0; i < 5; i++ {
		ans = (ans + dp[n-1][i]) % mod
	}

	return ans

}
