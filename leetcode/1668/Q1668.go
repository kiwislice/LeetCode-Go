package leetcode

import "strings"

func maxRepeating(sequence string, word string) int {
	pattern := word
	left := 0
	limit := len(sequence) / len(word)
	for i := 1; i <= limit; i++ {
		left = strings.Index(sequence, pattern)
		if left < 0 {
			return i - 1
		}
		pattern += word
	}
	return limit
}

func indexOf(sequence string, word string, from int) int {
	if from+len(word) > len(sequence) {
		return -1
	}

	for i := from; i <= len(sequence)-len(word); i++ {
		if sequence[i] != word[0] {
			continue
		}

		suc := true
		for j := 0; j < len(word); j++ {
			if sequence[i+j] != word[j] {
				suc = false
				break
			}
		}
		if suc {
			return i
		}
	}
	return -1
}
