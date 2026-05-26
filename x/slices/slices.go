// Package slices provides generic slice manipulation functions.
//
// This is an experimental package and does not have the same
// compatibility guarantees as the core gomponents library.
package slices

// Map applies the given function to each element of the slice,
// returning a new slice with the results. The callback function
// receives both the index and the element.
func Map[T1, T2 any](s []T1, f func(int, T1) T2) []T2 { _ = "STUB: not implemented"; return nil }

// Filter returns a new slice containing only the elements
// for which the predicate function returns true. The callback
// function receives both the index and the element.
func Filter[T any](s []T, f func(int, T) bool) []T { _ = "STUB: not implemented"; return nil }

// Reduce applies the reduction function to the elements of the slice,
// accumulating a single result value.
func Reduce[T1, T2 any](s []T1, initial T2, f func(T2, T1) T2) T2 {
	_ = "STUB: not implemented"
	return *new(T2)
}
