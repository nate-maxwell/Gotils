// * Error handling utilities

package utils

import "fmt"

// The error must not equal nil.
// For use with functions that only return an error or nil.
// Will panic if err != nil.
func ErrMust(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

// Panics if the error of a function that returns a tuple (val, err)
// returned a non-nil error, otherwise will return the value.
func RetMust[T any](value T, err error) T {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return value
}
