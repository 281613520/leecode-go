package contest303

import (
	"strconv"
)

func countExcellentPairs(nums []int, k int) (ans int64) {
	count := make(map[int]int)
	vis := make(map[int]bool)

	for _, v := range nums {
		if !vis[v] {
			count[countOne(v)]++
			vis[v] = true
		}
	}

	for cx, cxx := range count {
		for cy, cyy := range count {
			if cx+cy >= k {
				ans += int64(cxx) * int64(cyy)
			}
		}
	}

	return
}

func countOne(num int) int {
	result := strconv.FormatInt(int64(num), 2)

	ans := 0
	for _, v := range result {
		if v == '1' {
			ans++
		}
	}

	return ans
}
