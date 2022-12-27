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
