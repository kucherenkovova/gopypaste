package xmaps

// Keys returns a slice of all the keys in the map.
// The keys are in no particular order. Returns nil if the map is empty.
func Keys[K comparable, V any](m map[K]V) []K {
	if len(m) == 0 {
		return nil
	}

	keys := make([]K, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// Values returns a slice of all the values in the map.
// The values are in no particular order. Returns nil if the map is empty.
func Values[K comparable, V any](m map[K]V) []V {
	if len(m) == 0 {
		return nil
	}

	values := make([]V, 0, len(m))

	for _, v := range m {
		values = append(values, v)
	}

	return values
}
