package context292

import (
	"strings"
)

func largestGoodInteger(num string) string {

	for _, v := range "9876543210" {
		repeat := strings.Repeat(string(v), 3)
		if strings.Contains(num, repeat) {
			return repeat
		}
	}

	return ""
}
