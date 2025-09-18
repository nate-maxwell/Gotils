// * Error handling utilities

package errgo

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
func ValueOrPanic(err error) {
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
func PanicIfError[T any](value T, err error) T {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return value
}

// Boilerplate reduction for the err != nil check for prints.
func MessageIfError(msg string, err error) {
	if err != nil {
		formatted := fmt.Sprintf("%s: %s", msg, err)
		fmt.Println(formatted)
	}
}
