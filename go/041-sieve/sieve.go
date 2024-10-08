package sieve

import "slices"

func Sieve(limit int) []int {
	var marked []int
	var list []int

	for i := 2; i <= limit; i++ {
		// list = append(list, i)
		contains := slices.Contains(marked, i)

		if !contains {
			for j := i + i; j <= limit; j += i {
				// marked = append(marked, j)
				if !slices.Contains(marked, j) {
					marked = append(marked, j)
				}
			}
			list = append(list, i)
		}
	}

	return list
}
