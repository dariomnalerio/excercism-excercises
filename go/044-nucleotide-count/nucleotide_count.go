package dna

import "errors"

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	for _, nucleotide := range d {
		if _, ok := h[nucleotide]; ok {
			h[nucleotide]++
		} else {
			return h, errors.New("invalid nucleotide found")
		}
	}

	return h, nil
}
