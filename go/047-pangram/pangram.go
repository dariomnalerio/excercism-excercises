package pangram

import "unicode"

func isAlphabetLetter(r rune) bool {
	r = unicode.ToLower(r)
	return 'a' <= r && r <= 'z'
}

func IsPangram(input string) bool {
	alphabet := map[rune]bool{}

	for _, v := range input {
		if ok := isAlphabetLetter(v); ok {
			alphabet[unicode.ToLower(v)] = true
		}
	}

	var count int

	for _, exists := range alphabet {
		if exists {
			count++
		}
	}

	return count == 26
}
