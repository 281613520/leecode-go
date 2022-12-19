package main

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	l := 0
	r := 0
	max := 0
	for j := 0; j < n; j++ {
		for i := j - 1; i >= 0; i-- {
			if !dp[i][j] {
				if j-i == 1 {
					dp[i][j] = s[i] == s[j]
				} else {
					dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
				}
			}

			if dp[i][j] {
				if j-i > max {
					l = i
					r = j
					max = j - i
				}
			}

		}
	}

	return s[l : r+1]
}
