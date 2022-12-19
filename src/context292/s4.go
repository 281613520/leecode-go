package context292

func hasValidPath(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	if (m+n)%2 == 0 || grid[0][0] == ')' || grid[m-1][n-1] == '(' {
		return false
	}

	vis := make([][][]bool, m)
	for i := range vis {
		vis[i] = make([][]bool, n)
		for j := range vis[i] {
			vis[i][j] = make([]bool, m+n)
		}
	}
	var dfs func(int, int, int) bool
	dfs = func(x int, y int, c int) bool {
		if c > m+n-x-y-1 {
			return false
		}

		if x == m-1 && y == n-1 {
			return c == 1
		}

		if vis[x][y][c] { // 重复访问
			return false
		}
		vis[x][y][c] = true

		if grid[x][y] == '(' {
			c++
		} else {
			c--
			if c < 0 {
				return false
			}
		}

		return x < m-1 && dfs(x+1, y, c) || y < n-1 && dfs(x, y+1, c)
	}

	return dfs(0, 0, 0)
}
