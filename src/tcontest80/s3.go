package tcontest80

import (
	"strings"
)

func matchReplacement(s string, sub string, mappings [][]byte) bool {
	if strings.Contains(s, sub) {
		return true
	}

	charMap := make(map[byte]map[byte]bool, 0)

	for _, v := range mappings {
		key, value := v[0], v[1]
		charMap[key] = make(map[byte]bool, 0)
		charMap[key][value] = true
	}

	m, n := len(s), len(sub)
	count := 0
	for i := m - n; i < m; i++ {
		tmp := s[i-n : i]

		for j := range tmp {
			if tmp[j] == sub[j] || charMap[sub[j]][tmp[j]] {
				count++
			}
		}

		if count == n {
			return true
		}

		count = 0

	}

	return false

}
