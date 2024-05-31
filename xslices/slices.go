package xslices

// Unique returns a new slice containing only the unique elements of the input slice.
// Only first occurrence of each element is preserved.
func Unique[T comparable](list []T) []T {
	if list == nil {
		return nil
	}

	seen := make(map[T]struct{}, len(list))
	results := make([]T, 0, len(list))

	for _, element := range list {
		if _, ok := seen[element]; ok {
			continue
		}

		seen[element] = struct{}{}
		results = append(results, element)
	}

	return results
}
