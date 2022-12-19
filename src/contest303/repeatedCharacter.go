package contest303

func repeatedCharacter(s string) byte {
	wordMap := make(map[byte]bool)

	for _, v := range s {
		if wordMap[byte(v)] {
			return byte(v)
		} else {
			wordMap[byte(v)] = true
		}
	}

	return byte(0)
}
