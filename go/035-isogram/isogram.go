package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	set := make(map[rune]bool)
	for _, v := range strings.ToLower(word) {
		if !unicode.IsLetter(v) {
			continue
		}
		if _, ok := set[v]; ok {
			return false
		}
		set[v] = true
	}
	return true
}
