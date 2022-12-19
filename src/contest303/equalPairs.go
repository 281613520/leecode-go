package contest303

func equalPairs(grid [][]int) int {
	ans := 0
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			flag := true
			for k := 0; k < m; k++ {
				if grid[i][k] != grid[k][j] {
					flag = false
				}
			}

			if flag {
				ans++
			}
		}
	}

	return ans
}
