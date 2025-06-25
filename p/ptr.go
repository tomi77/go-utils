package p

// Ptr returns a pointer to the value v.
// It is a convenience function to avoid writing &v.
// Example usage:
//
//	p.Ptr(42) // returns *int pointing to 42
//	p.Ptr("hello") // returns *string pointing to "hello"
//	p.Ptr([]int{1, 2, 3}) // returns *[]int pointing to []int{1, 2, 3}
//	p.Ptr(map[string]int{"a": 1, "b": 2}) // returns *map[string]int pointing to map[string]int{"a": 1, "b
func Ptr[T any](v T) *T {
	return &v
}
