package dp

func largestVariance(s string) int {
	ans := 0
	// 1.枚举两个字母
	// 2.第一个字母就+1，第二个字母就-1
	for a := 'a'; a <= 'z'; a++ {
		for b := 'a'; b <= 'z'; b++ {
			if a == b {
				continue
			}
			diff, diffB := 0, -len(s)
			for _, s := range s {
				if s == a {
					diff++
					diffB++
				} else if s == b {
					diff--
					diffB = diff
					diff = max(0, diff)
				}
				ans = max(diffB, ans)
			}
		}
	}
	return ans
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
