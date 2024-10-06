package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

var wordRegex = regexp.MustCompile(`\b[\w']+\b`)

func WordCount(phrase string) Frequency {
	freq := make(Frequency)

	phrase = strings.ToLower(phrase)

	words := wordRegex.FindAllString(phrase, -1)

	for _, word := range words {
		freq[word]++
	}

	return freq
}
