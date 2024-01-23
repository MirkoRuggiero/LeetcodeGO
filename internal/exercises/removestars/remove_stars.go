package removestars

import (
	"LeetcodeGO/internal/utils/ds"
	"strings"
)

func RemoveStars(args ...interface{}) interface{} {
	var input = args[0].(string)

	return removeStars(input)
}

func removeStars(s string) string {
	stack := ds.NewStack[rune]()

	for _, c := range s {
		if c != '*' {
			stack.Push(c)
		} else {
			stack.Pop()
		}
	}
	var result strings.Builder

	for stack.Length() > 0 {
		r := stack.Pop()
		result.WriteRune(r)
	}

	runes := []rune(result.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
