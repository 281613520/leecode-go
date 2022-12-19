package main

func firstDayBeenInAllRooms(nextVisit []int) int {
	mod := 1000000007
	n := len(nextVisit)
	dp := make([]int, n)
	dp[0] = 0
	dp[1] = 2

	for i := 2; i < n; i++ {
		dp[i] = (2*dp[i-1] - dp[nextVisit[i-1]] + 2) % mod
		if dp[i] < 0 {
			dp[i] += mod
		}
	}

	return dp[n-1]
}
