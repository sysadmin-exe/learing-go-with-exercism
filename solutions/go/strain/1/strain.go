package strain

// Keep returns a new slice containing elements that satisfy the predicate.
func Keep[T any](s []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range s {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Discard returns a new slice excluding elements that satisfy the predicate.
func Discard[T any](s []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range s {
		if !predicate(v) {
			result = append(result, v)
		}
	}
	return result
}
