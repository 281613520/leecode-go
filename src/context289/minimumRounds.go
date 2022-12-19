package context289

func MinimumRounds(tasks []int) int {
	numsMap := make(map[int]int, 0)
	for _, v := range tasks {
		numsMap[v]++
	}

	ans := 0

	for _, v := range numsMap {
		if v%3 == 2 {
			ans += v/3 + 1
		} else if v%3 == 0 {
			ans += v / 3
		} else if v%3 == 1 {
			n := v / 3
			if n > 0 && (v-(n-1)*3)%2 == 0 {
				ans += v/3 + 1
			} else {
				return -1
			}
		} else {
			return -1
		}
	}

	return ans
}
