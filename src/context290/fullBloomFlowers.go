package context290

import (
	"sort"
)

func fullBloomFlowers(flowers [][]int, persons []int) []int {
	ans := make([]int, len(persons))

	n := len(flowers)
	op := make([]int, n)
	cl := make([]int, n)

	for i, v := range flowers {
		op[i] = v[0]
		cl[i] = v[1]
	}
	sort.Ints(op)
	sort.Ints(cl)

	for i, p := range persons {
		opNums := sort.SearchInts(op, p+1)
		clNums := sort.SearchInts(cl, p)
		ans[i] = opNums - clNums
	}

	return ans
}
