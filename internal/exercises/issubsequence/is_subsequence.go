package issubsequence

import (
	"strings"
)

func IsSubsequence(args ...interface{}) interface{} {
	var s = args[0].(string)
	var t = args[1].(string)

	return isSubsequence(s, t)
}

func isSubsequence(s, t string) bool {
	currentIndex := -1

	for _, sChar := range s {
		currentIndex = findCharInStringStartingFromPosition(sChar, t, currentIndex)
		if currentIndex == -1 {
			return false
		}
	}
	return true
}

func findCharInStringStartingFromPosition(sChar rune, t string, currentIndex int) int {
	if currentIndex == -1 {
		return strings.IndexRune(t, sChar)
	}
	for i := currentIndex + 1; i < len(t); i++ {
		if rune(t[i]) == sChar {
			return i
		}
	}
	return -1
}
