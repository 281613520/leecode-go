package graph

func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	var dfs func(image [][]int, x int, y int, sourceColor int, color int)

	dfs = func(image [][]int, x int, y int, sourceColor int, color int) {
		if x < 0 || y < 0 || x >= len(image) || y >= len(image[0]) {
			return
		}
		if image[x][y] == sourceColor {
			image[x][y] = color

			dfs(image, x+1, y, sourceColor, color)
			dfs(image, x-1, y, sourceColor, color)
			dfs(image, x, y+1, sourceColor, color)
			dfs(image, x, y-1, sourceColor, color)
		}
	}

	source := image[sr][sc]

	if source != color {
		dfs(image, sr, sc, source, color)
	}

	return image
}

func numIslands(grid [][]byte) int {
	var dfs func(grid [][]byte, x int, y int)

	dfs = func(grid [][]byte, x int, y int) {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] != '1' {
			return
		}
		grid[x][y] = '2'
		dfs(grid, x+1, y)
		dfs(grid, x-1, y)
		dfs(grid, x, y+1)
		dfs(grid, x, y-1)
	}

	ans := 0

	for i, v1 := range grid {
		for j, v2 := range v1 {
			if v2 == '1' {
				dfs(grid, i, j)
				ans++
			}
		}
	}

	return ans
}

func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	var dfs func(grid [][]int, x int, y int) int

	dfs = func(grid [][]int, x int, y int) int {

		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] != 1 {
			return 0
		}

		tmp := 1
		grid[x][y] = 2
		tmp += dfs(grid, x+1, y)
		tmp += dfs(grid, x-1, y)
		tmp += dfs(grid, x, y+1)
		tmp += dfs(grid, x, y-1)

		return tmp
	}

	for i, v1 := range grid {
		for j, v2 := range v1 {
			if v2 == 1 {
				ans = max(dfs(grid, i, j), ans)

			}
		}
	}

	return ans
}

func max(mianji int, ans int) int {
	if mianji > ans {
		return mianji
	}

	return ans
}

func closedIsland(grid [][]int) int {
	nums := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] == 0 {
				if dfs(grid, i, j, true) {
					nums++
				}
			}
		}
	}
	return nums
}

func dfs(grid [][]int, r, c int, result bool) bool {
	if r == 0 || c == 0 || r == len(grid)-1 || c == len(grid[r])-1 {
		return false
	}
	grid[r][c] = 1
	if r-1 >= 0 && grid[r-1][c] == 0 {
		// result放到后边，避免提前短路
		result = dfs(grid, r-1, c, result) && result
	}
	if r+1 < len(grid) && grid[r+1][c] == 0 {
		result = dfs(grid, r+1, c, result) && result
	}
	if c-1 >= 0 && grid[r][c-1] == 0 {
		result = dfs(grid, r, c-1, result) && result
	}
	if c+1 < len(grid[r]) && grid[r][c+1] == 0 {
		result = dfs(grid, r, c+1, result) && result
	}
	return result
}

func numEnclaves(grid [][]int) int {
	ans := 0
	var dfs func(grid [][]int, x int, y int)

	dfs = func(grid [][]int, x int, y int) {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == 0 {
			return
		}
		grid[x][y] = 0

		dfs(grid, x+1, y)
		dfs(grid, x-1, y)
		dfs(grid, x, y+1)
		dfs(grid, x, y-1)
	}

	for i, v := range grid {
		for j := range v {
			if i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[0])-1 {
				dfs(grid, i, j)
			}
		}
	}

	for i, v := range grid {
		for j := range v {
			if grid[i][j] > 0 {
				ans++
			}
		}
	}
	return ans
}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	flag := false
	ans := 0

	var dfs func(grid [][]int, x int, y int)

	dfs = func(grid [][]int, x int, y int) {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == 0 {
			return
		}

		if grid1[x][y] != 1 {
			flag = false
		}

		grid[x][y] = 0

		dfs(grid, x+1, y)
		dfs(grid, x-1, y)
		dfs(grid, x, y+1)
		dfs(grid, x, y-1)
	}

	for i, v := range grid2 {
		for j := range v {
			if grid2[i][j] > 0 {
				flag = true
				dfs(grid2, i, j)
				if flag {
					ans++
				}
			}
		}
	}
	return ans

}
