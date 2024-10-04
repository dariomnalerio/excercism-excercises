package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("string sequences must be of equal length")
	}
	var distance int
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
