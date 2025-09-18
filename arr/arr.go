package arr

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/constraints"
)

// FindIndex returns the index of an element in a slice.
// If the element is not found, it returns -1.
func FindIndex[T comparable](arr []T, target T) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

// Unique removes duplicate elements from a slice and preserves order.
// Works for any comparable type (ints, strings, floats, structs with comparable
// fields).
func Unique[T comparable](arr []T) []T {
	seen := make(map[T]struct{})
	out := make([]T, 0, len(arr))

	for _, v := range arr {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			out = append(out, v)
		}
	}
	return out
}

// Shuffle shuffles a slice of any element type in place.
func Shuffle[T any](arr []T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(arr) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// Reverse reverses a slice of any element type in place.
func Reverse[T any](arr []T) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// Sum computes the sum of all elements in a slice of numeric types.
func Sum[T constraints.Integer | constraints.Float](arr []T) T {
	var total T
	for _, v := range arr {
		total += v
	}
	return total
}

// -------To String-------------------------------------------------------------

// Returns a comma separated string of integers.
func ConvertIntSliceToString(input []int) string {
	var output []string
	for _, i := range input {
		output = append(output, strconv.Itoa(i))
	}
	return strings.Join(output, ",")
}

// Returns a comma separated string of float32.
func ConvertF32SliceToString(input []float32) string {
	var output []string
	for _, i := range input {
		output = append(output, strconv.FormatFloat(float64(i), 'f', -1, 32))
	}
	return strings.Join(output, ",")
}

// Returns a comma separated string of float64.
func ConvertF64SliceToString(input []float64) string {
	var output []string
	for _, i := range input {
		output = append(output, strconv.FormatFloat(i, 'f', -1, 64))
	}
	return strings.Join(output, ",")
}
