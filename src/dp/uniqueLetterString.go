package dp

func uniqueLetterString(s string) int {
	locationMap := make([][]int, 26)

	for i := range locationMap {
		locationMap[i] = make([]int, 0)
	}

	for i, v := range s {
		idx := v - 'A'
		locationMap[idx] = append(locationMap[idx], i)
	}

	ans := 0

	for _, idxs := range locationMap {
		for i := range idxs {
			prev := -1
			next := len(s)
			if i > 0 {
				prev = idxs[i-1]
			}

			if i < len(idxs)-1 {
				next = idxs[i+1]
			}

			ans += (idxs[i] - prev) * (next - idxs[i])
		}
	}

	return ans
}
