package src

func minOperations(boxes string) []int {
	left := int(boxes[0] - '0')
	right := 0
	operations := 0
	n := len(boxes)
	for i := 1; i < n; i++ {
		if boxes[i] == '1' {
			right++
			operations += i
		}
	}

	ans := make([]int, n)
	ans[0] = operations
	for i := 1; i < n; i++ {
		operations = operations + left - right
		if boxes[i] == '1' {
			left++
			right--
		}
		ans[i] = operations
	}

	return ans
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	graphMap := make([][]int, n)
	visited := make([]bool, n)

	for _, v := range edges {
		start := v[0]
		end := v[1]
		graphMap[start] = append(graphMap[start], end)
		graphMap[end] = append(graphMap[end], start)
	}

	var dfs func(start int, end int)
	result := false

	dfs = func(start int, end int) {

		if start == end {
			result = true
			return
		}
		if visited[start] {
			return
		}

		visited[start] = true

		next := graphMap[start]

		for _, v := range next {
			dfs(v, end)
		}
	}

	dfs(source, destination)

	return result
}
