package main

import "sort"

func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] < properties[j][0] {
			return true
		} else if properties[i][0] == properties[j][0] {
			if properties[i][0] > properties[j][0] {
				return false
			} else {
				return true
			}
		}

		return false
	})

	ans := 0

	return ans
}
