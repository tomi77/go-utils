# Go utils

## packages

### p

Pointers

#### Ptr

Ptr returns a pointer to the value v.
It is a convenience function to avoid writing &v.

### s

Slices

#### Map

Map applies the function fn to each element of the slice m and returns a new slice containing the results.
If m is empty, it returns nil.

#### MapInPlace

MapInPlace applies the function fn to each element of the slice m in place.
It modifies the original slice and does not return a new slice.
