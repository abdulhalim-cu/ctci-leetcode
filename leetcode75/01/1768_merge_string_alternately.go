package _1

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	lenW1, lenW2 := len(word1), len(word2)
	minLen := lenW1
	if lenW2 < lenW1 {
		minLen = lenW2
	}

	var builder strings.Builder
	for i := 0; i < minLen; i++ {
		builder.WriteByte(word1[i])
		builder.WriteByte(word2[i])
	}

	if lenW1 > lenW2 {
		builder.WriteString(word1[minLen:])
	} else if lenW2 > lenW1 {
		builder.WriteString(word2[minLen:])
	}

	return builder.String()
}
