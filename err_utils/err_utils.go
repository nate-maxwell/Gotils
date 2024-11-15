// * Error handling utilities

package err_utils

import "fmt"

// The error must not equal nil.
// For use with functions that only return an error or nil.
// Will panic if err != nil.
//
// Instead of writing:
// >>> err := someFunc(args...)
// >>> if err != nil {}
//
// You can write:
// >>> utils.ErrMust(someFunc(args...))
// and it will panic if the error is not nil.
func ErrMust(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

// Panics if the error of a function that returns a tuple (val, err)
// returned a non-nil error, otherwise will return the value.
//
// Instead of writing:
// >>> foo, err := someFunc(args...)
// >>> if err != nil {}
//
// You can write:
// >>> foo := utils.RetMust(someFunc(args...))
// and it will panic if the error is not nil.
func RetMust[T any](value T, err error) T {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return value
}
