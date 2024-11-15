package arr

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Removes duplicate entries from a slice of strings.
//
// Args:
//
//	strSlice([]string): The slice of strings to minify.
//
// Returns:
//
//	[]string: The minified slice of strings.
func RemoveDuplicatesFromSlice(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// -------------------------------------------------- Shuffles --------------------------------------------------

// Shuffles a string slice in place.
func ShuffleStringSlice(arr []string) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(arr) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// Shuffles a string slice in place.
func ShuffleIntSlice(arr []int) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(arr) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// Shuffles a string slice in place.
func ShuffleF32Slice(arr []float32) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(arr) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// Shuffles a string slice in place.
func ShuffleF64Slice(arr []float64) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(arr) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// -------------------------------------------------- Reverse --------------------------------------------------

// Reverses the order of a slice of string.
func ReverseStringSlice(a []string) []string {
	for i := len(a)/2 - 1; i >= 0; i-- {
		pos := len(a) - 1 - i
		a[i], a[pos] = a[pos], a[i]
	}
	return a
}

// Reverses the order of a slice of integers.
func ReverseIntSlice(a []int) []int {
	for i := len(a)/2 - 1; i >= 0; i-- {
		pos := len(a) - 1 - i
		a[i], a[pos] = a[pos], a[i]
	}
	return a
}

// Reverses the order of a slice of float32.
func ReverseF32Slice(a []float32) []float32 {
	for i := len(a)/2 - 1; i >= 0; i-- {
		pos := len(a) - 1 - i
		a[i], a[pos] = a[pos], a[i]
	}
	return a
}

// Reverses the order of a slice of float64.
func ReverseF64Slice(a []float64) []float64 {
	for i := len(a)/2 - 1; i >= 0; i-- {
		pos := len(a) - 1 - i
		a[i], a[pos] = a[pos], a[i]
	}
	return a
}

// -------------------------------------------------- Sum --------------------------------------------------

// Sum the contents of a slice of integers.
func SumIntegerSlice(arr []int) int {
	sum := 0
	for _, i := range arr {
		sum += i
	}
	return sum
}

// Sum the contents of a slice of float32.
func SumF32Slice(arr []float32) float32 {
	sum := float32(0)
	for _, i := range arr {
		sum += i
	}
	return sum
}

// Sum the contents of a slice of float64.
func SumF64lice(arr []float64) float64 {
	sum := float64(0)
	for _, i := range arr {
		sum += i
	}
	return sum
}

// -------------------------------------------------- To String --------------------------------------------------

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
