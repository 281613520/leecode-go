package context289

type numCount struct {
	num2 int
	num5 int
}

// 路径中2和5中数量少 的最大值
func MaxTrailingZeros(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])

	ans := 0

	preSumInRow := make([][]numCount, 0)
	preSumInCol := make([][]numCount, 0)

	for i := 0; i < row; i++ {
		preSumInRow[i] = make([]numCount, col)
		preSumInCol[i] = make([]numCount, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			num2, num5 := getNum(grid[i][j])
			if j == 0 {
				preSumInRow[i][j] = numCount{
					num2: num2,
					num5: num5,
				}
			} else {
				pre := preSumInRow[i][j-1]
				preSumInRow[i][j] = numCount{
					num2: num2 + pre.num2,
					num5: num5 + pre.num5,
				}
			}
		}
	}

	for j := 0; j < col; j++ {
		for i := 0; i < row; i++ {
			num2, num5 := getNum(grid[i][j])
			if i == 0 {
				preSumInCol[i][j] = numCount{
					num2: num2,
					num5: num5,
				}
			} else {
				pre := preSumInRow[i-1][j]
				preSumInRow[i][j] = numCount{
					num2: num2 + pre.num2,
					num5: num5 + pre.num5,
				}
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			//四个方向
			// 从左边出发，到上边结束
			ans = max(ans, min(preSumInRow[i][j].num2+preSumInCol[i-1][j].num2, preSumInRow[i][j].num5+preSumInCol[i-1][j].num5))
			// 从左边出发，到下边结束
			ans = max(ans, min(preSumInRow[i][j].num2+preSumInCol[row][j].num2-preSumInCol[i][j].num2, preSumInRow[i][j].num5+preSumInCol[row][j].num5-preSumInCol[i][j].num5))
			// 从右边出发，到上边结束
			ans = max(ans, min(preSumInRow[col][j].num2-preSumInRow[i][j].num2+preSumInCol[i][j].num2, preSumInRow[col][j].num5-preSumInRow[i][j].num5+preSumInCol[i][j].num5))
			// 从右边出发，到下边结束
			ans = max(ans, min(preSumInRow[col][j].num2-preSumInRow[i][j].num2+preSumInCol[row][j].num2-preSumInCol[i][j].num2, preSumInRow[col][j].num5-preSumInRow[i][j].num5+preSumInCol[row][j].num5-preSumInCol[i][j].num5))
		}
	}

	return ans
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x int, y int) int {
	if x > y {

		return x
	}
	return y
}

func getNum(i int) (int, int) {
	num2 := 0
	num5 := 0

	for cur := i; cur%2 == 0 && cur > 0; cur = cur / 2 {
		num2++
	}

	for cur := i; cur%5 == 0 && cur > 0; cur = cur / 5 {
		num5++
	}

	return num2, num5
}
