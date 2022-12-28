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

func countHomogenous(s string) int {
	const mod int = 1e9 + 7
	cnt := 0
	prev := rune(s[0])
	ans := 0

	for _, v := range s {
		if v == prev {
			cnt++
		} else {
			prev = v

			ans += cnt * (cnt + 1) / 2 % mod

			cnt = 1
		}
	}
	ans += (cnt + 1) * cnt / 2 % mod
	return ans

}

func dividePlayers(skill []int) int64 {
	sum := 0

	numMap := make(map[int]int)

	for _, v := range skill {
		sum += v
		numMap[v]++
	}

	nums := len(skill) / 2

	if sum%nums != 0 {
		return -1
	}

	avg := sum / nums

	ans := int64(0)

	for _, v := range skill {
		cur := avg - v

		if numMap[cur] == numMap[v] {
			ans += int64(cur * v)
		} else {
			return -1
		}

	}

	return ans / 2
}

func minimumMoves(s string) int {
	ans := 0
	covered := -1

	for i, v := range s {
		if v == 'X' && i > covered {
			ans++
			covered = i + 2
		}
	}

	return ans
}

func minimumLength(s string) int {
	i, j := 0, len(s)

	for i < j && s[i] == s[j] {
		c := s[i]

		for i <= j && s[i] == c {
			i++
		}

		for i <= j && s[j] == c {
			j--
		}
	}

	return j - i + 1
}
