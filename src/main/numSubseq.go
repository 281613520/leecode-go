package main

import "sort"

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = dp[i-1] * 2 % (1e9 + 7)
	}
	l, r, ans := 0, len(nums)-1, 0
	for l <= r && 2*l <= len(nums) {
		tmp := nums[l] + nums[r]
		if tmp <= target {
			ans = (ans + dp[r-l]) % (1e9 + 7)
			l++
		} else {
			r--
		}
	}

	return ans
}
