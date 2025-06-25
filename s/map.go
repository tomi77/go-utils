package s

// Map applies the function fn to each element of the slice m and returns a new slice containing the results.
// If m is empty, it returns nil.
// Example usage:
//
//	s.Map([]int{1, 2, 3}, func(i int) int {
//	    return i * 2
//	}) // returns []int{2, 4, 6}
//	s.Map([]string{"a", "b", "c"}, func(s string) string {
//	    return s + "!"
//	}) // returns []string{"a!", "b!", "c!"}
//	s.Map([]float64{1.1, 2.2, 3.3}, func(f float64) float64 {
//	    return f + 1.0
//	}) // returns []float64{2.1, 3.2, 4.3}
//	s.Map([]bool{true, false}, func(b bool) bool {
//	    return !b
//	}) // returns []bool{false, true}
//	s.Map([]string{}, func(s string) string {
//	    return s + "!"
//	}) // returns nil
//	s.Map([]int{1, 2, 3}, nil) // returns nil
//	s.Map(nil, func(i int) int {
//	    return i * 2
//	}) // returns nil
func Map[K any, V any](m []K, fn func(K) V) []V {
	if m == nil || fn == nil {
		return nil
	}
	r := make([]V, len(m))
	for i, v := range m {
		r[i] = fn(v)
	}
	return r
}

// MapInPlace applies the function fn to each element of the slice m in place.
// It modifies the original slice and does not return a new slice.
// If m is nil or fn is nil, it does nothing.
// Example usage:
//
//	s.MapInPlace([]int{1, 2, 3}, func(i int) int {
//	    return i * 2
//	}) // modifies the original slice to []int{2, 4, 6}
//	s.MapInPlace([]string{"a", "b", "c"}, func(s string) string {
//	    return s + "!"
//	}) // modifies the original slice to []string{"a!", "b!", "c!"}
//	s.MapInPlace([]float64{1.1, 2.2, 3.3}, func(f float64) float64 {
//	    return f + 1.0
//	}) // modifies the original slice to []float64{2.1, 3.2, 4.3}
//	s.MapInPlace([]bool{true, false}, func(b bool) bool {
//	    return !b
//	}) // modifies the original slice to []bool{false, true}
func MapInPlace[K any](m []K, fn func(K) K) {
	if m == nil || fn == nil {
		return
	}
	for i, v := range m {
		m[i] = fn(v)
	}
}
