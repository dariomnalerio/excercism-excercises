package collatzconjecture

import "fmt"

func Step(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return n*3 + 1
}

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("input must be a positive integer")
	}
	var steps int
	aux := n
	for aux != 1 {
		aux = Step(aux)
		steps++
	}
	return steps, nil
}
