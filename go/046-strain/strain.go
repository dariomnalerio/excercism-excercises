package strain

// Keep returns a new slice containing elements that satisfy the predicate
func Keep[T any](collection []T, predicate func(T) bool) []T {
	var result []T

	for _, v := range collection {
		if ok := predicate(v); ok {
			result = append(result, v)
		}
	}

	return result
}

// Discard returns a new slice containing elements that do not satisfy the predicate
func Discard[T any](collection []T, predicate func(T) bool) []T {
	var result []T

	for _, v := range collection {
		if ok := predicate(v); !ok {
			result = append(result, v)
		}
	}

	return result
}
