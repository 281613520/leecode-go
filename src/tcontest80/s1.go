package tcontest80

import "strings"

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}

	first, second, third, forth, fifth := false, false, false, false, true
	lastChar := '0'
	template := "!@#$%^&*()-+"

	for _, c := range password {
		if c >= 'a' && c <= 'z' {
			first = true
		}

		if c >= 'A' && c <= 'Z' {
			second = true
		}

		if c >= '0' && c <= '9' {
			third = true
		}

		if strings.Contains(template, string(c)) {
			forth = true
		}

		if lastChar == c {
			fifth = false
		}

		lastChar = c
	}

	return first && second && forth && third && fifth
}
