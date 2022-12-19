package tcontest80

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	ans := make([]int, len(spells))

	sort.Ints(potions)

	for i, v := range spells {
		location := sort.SearchInts(potions, (int(success)-1)/v+1)
		ans[i] = len(potions) - location
	}

	return ans
}
