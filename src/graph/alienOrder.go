package graph

func AlienOrder(words []string) string {
	g := map[byte][]byte{}
	inDegree := map[byte]int{}

	for _, c := range words[0] {
		inDegree[byte(c)] = 0
	}

	for i := 1; i < len(words); i++ {
		first := words[i-1]
		second := words[i]
		minLen := min(len(first), len(second))

		for _, v := range second {
			inDegree[byte(v)] += 0
		}
		flag := true
		for j := 0; j < minLen; j++ {
			if first[j] != second[j] {
				g[byte(first[j])] = append(g[byte(first[j])], byte(second[j]))
				inDegree[byte(second[j])]++
				flag = false
				break
			}
		}

		if len(first) > len(second) && flag {
			return ""
		}
	}
	ans := make([]byte, 0)
	queue := make([]byte, 0)
	for k, v := range inDegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}
	count := 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		ans = append(ans, cur)
		for _, v := range g[cur] {
			if inDegree[v]--; inDegree[v] == 0 {
				count++
				queue = append(queue, v)
			}
		}
	}

	if count == len(inDegree) {
		return string(ans)
	}

	return ""
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
