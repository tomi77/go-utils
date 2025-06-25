# Go utils

## p

Pointers

```go
Ptr[T any](v T) *T
```

Returns a pointer to the value v.
It is a convenience function to avoid writing &v.

## s

Slices

```go
Map[K comparable, V any](m []K, fn func(K) V) []V
```

Applies the function fn to each element of the slice m and returns a new slice containing the results.
If m is empty, it returns nil.

```go
MapInPlace[K comparable](m []K, fn func(K) K)
```

Applies the function fn to each element of the slice m in place.
It modifies the original slice and does not return a new slice.

```go
Unique[S ~[]E, E comparable](m S) S
```

Returns a new slice containing only the unique elements from the input slice.
It preserves the order of the first occurrence of each element.

```go
UniqueFunc[S ~[]E, E comparable](m S, fn func(E) E) S
```

Returns a new slice containing only the unique elements from the input slice, based on the transformation function fn applied to each element.
It preserves the order of the first occurrence of each transformed element.