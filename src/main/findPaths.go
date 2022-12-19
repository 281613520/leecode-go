package main

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {

	const mod = 7 + 1e9
	cache := make([][][]int, m)

	for i := 0; i < m; i++ {
		cache[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			cache[i][j] = make([]int, maxMove+1)
			for k := 0; k <= maxMove; k++ {
				cache[i][j][k] = -1
			}
		}
	}

	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(int, int, int) int

	dfs = func(row int, col int, k int) int {
		if row < 0 || row == m || col < 0 || col == n {
			return 1
		}
		if k == maxMove {
			return 0
		}

		if cache[row][col][k] != -1 {
			return cache[row][col][k]
		}

		sum := 0

		for _, v := range dir {
			x := v[0]
			y := v[1]
			sum += dfs(row+x, col+y, k+1)
			sum = sum % mod
		}

		cache[row][col][k] = sum
		return sum

	}

	return dfs(startRow, startColumn, 0)
}
