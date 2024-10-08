package armstrong

import "math"

func IsNumber(n int) bool {
	aux := n
	var slice []int

	for aux > 0 {
		slice = append(slice, aux%10)
		aux /= 10
	}

	var result int
	for _, v := range slice {
		result += int(math.Pow(float64(v), float64(len(slice))))
	}

	return result == n
}
