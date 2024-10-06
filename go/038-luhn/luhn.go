package luhn

import (
	"strings"
	"unicode"
)

// Valid checks if the provided ID is a valid Luhn number.
func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) < 2 {
		return false
	}

	var sum int
	double := false

	for i := len(id) - 1; i >= 0; i-- {
		char := id[i]

		if !unicode.IsDigit(rune(char)) {
			return false
		}

		num := int(char - '0')

		if double {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}

		sum += num
		double = !double
	}

	return sum%10 == 0
}
