package s

// Unique returns a new slice containing only the unique elements from the input slice.
// It preserves the order of the first occurrence of each element.
// If the input slice is nil, it returns nil.
// Example usage:
//
//	s.Unique([]int{1, 2, 3, 1, 2}) // returns []int{1, 2, 3}
//	s.Unique([]string{"a", "b", "c", "a"}) // returns []string{"a", "b", "c"}
//	s.Unique([]int{1, 2, 3, 1, 2, 3}) // returns []int{1, 2, 3}
func Unique[S ~[]E, E comparable](m S) S {
	if m == nil {
		return nil
	}

	seen := make(map[E]struct{}, len(m))
	result := make(S, 0, len(m))

	for _, v := range m {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}

// UniqueFunc returns a new slice containing only the unique elements from the input slice,
// based on the transformation function fn applied to each element.
// It preserves the order of the first occurrence of each transformed element.
// If the input slice is nil or the function fn is nil, it returns nil.
// Example usage:
//
//	s.UniqueFunc([]string{"apple", "banana", "applE", "orange"}, func(s string) string { return s[:3] }) // returns []string{"apple", "banana", "orange"}
//	s.UniqueFunc([]string{"a", "b", "c", "A"}, strings.ToLower) // returns []string{"a", "b", "c"}
func UniqueFunc[S ~[]E, E comparable](m S, fn func(E) E) S {
	if m == nil || fn == nil {
		return nil
	}

	seen := make(map[E]struct{}, len(m))
	result := make(S, 0, len(m))

	for _, v := range m {
		transformed := fn(v)
		if _, exists := seen[transformed]; !exists {
			seen[transformed] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}
