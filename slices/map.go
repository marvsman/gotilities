// Package slices provides utility functions for working with slices / arrays.
// This should mimic some parts of javascript's array methods like map, reduce, filter, etc.
package slices

type NumericTypes interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// Map applies the function fn to each element of the input slice ts and returns a new slice containing the results.
func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

// Sum calculates the sum of the values obtained by applying the function fn to each element of the input slice ts.
func Sum[T, V NumericTypes](ts []T, fn func(T) V) V {
	var result V
	for _, t := range ts {
		result += fn(t)
	}
	return result
}

// Avg calculates the average of the values obtained by applying the function fn to each element of the input slice ts.
func Avg[T, V NumericTypes](ts []T, fn func(T) V) V {
	if len(ts) == 0 {
		var zero V
		return zero
	}
	return Sum(ts, fn) / V(len(ts))
}

// Min returns the minimum value obtained by applying the function fn to each element of the input slice ts.
func Min[T, V NumericTypes](ts []T, fn func(T) V) V {
	if len(ts) == 0 {
		var zero V
		return zero
	}
	m := fn(ts[0])
	for _, t := range ts[1:] {
		v := fn(t)
		if v < m {
			m = v
		}
	}
	return m
}

// Max returns the maximum value obtained by applying the function fn to each element of the input slice ts.
func Max[T, V NumericTypes](ts []T, fn func(T) V) V {
	if len(ts) == 0 {
		var zero V
		return zero
	}
	m := fn(ts[0])
	for _, t := range ts[1:] {
		v := fn(t)
		if v > m {
			m = v
		}
	}
	return m
}

// Reduce reduces the input slice ts to a single value by applying the function fn, starting with the initial value.
func Reduce[T, V any](ts []T, fn func(acc V, t T) V, initial V) V {
	acc := initial
	for _, t := range ts {
		acc = fn(acc, t)
	}
	return acc
}
