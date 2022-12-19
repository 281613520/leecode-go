package main

import "strings"

func reversePrefix(word string, ch byte) string {
	j := strings.IndexByte(word, ch)
	if j < 0 {
		return word
	}
	t := []byte(word)
	for i := 0; i < j; i++ {
		t[i], t[j] = t[j], t[i]
		j--
	}
	return string(t)

}
