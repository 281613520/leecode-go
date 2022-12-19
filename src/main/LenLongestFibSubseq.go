package main

func lenLongestFibSubseq(arr []int) int {
	n := len(arr)

	location := make(map[int]int)

	for i, a := range arr {
		location[a] = i
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// arr[k]  arr[j] arr[i]
	// dp[j][i] = dp[k][j] +1
	// dp[j][i] = 3
	ans := 0
	for i := range arr {
		for j := n - 1; j >= 0 && arr[j]*2 > arr[i]; j-- {
			if k, ok := location[arr[i]-arr[j]]; ok {
				dp[j][i] = max(3, dp[k][j]+1)
				ans = max(ans, dp[j][i])
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
