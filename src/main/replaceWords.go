package main

import (
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	dictionarySet := make(map[string]bool)

	for _, w := range dictionary {
		dictionarySet[w] = true
	}

	words := strings.Split(sentence, " ")

	for i, w := range words {
		for j := 0; j < len(w); j++ {
			if dictionarySet[w[:j]] {
				words[i] = w[:j]
				break
			}
		}
	}

	return strings.Join(words, " ")

}
