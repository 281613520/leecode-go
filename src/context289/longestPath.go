package context289

func LongestPath(parent []int, s string) int {
	ans := 0
	var dfs func(node int) (maxLen int)

	martix := make([][]int, len(parent))
	for i, v := range parent {
		if v == -1 {
			continue
		} else {
			if martix[v] == nil {
				martix[v] = make([]int, 0)
			}

			martix[v] = append(martix[v], i)
		}
	}

	dfs = func(node int) (maxLen int) {
		for _, v := range martix[node] {
			len := dfs(v) + 1
			if s[v] != s[node] {
				ans = max(maxLen+len, ans)
				maxLen = max(maxLen, len)
			}
		}

		return
	}

	dfs(0)
	return ans + 1
}
