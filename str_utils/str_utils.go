package str_utils

import (
	"regexp"
	"strings"

	"github.com/chigopher/pathlib"
)

// Converts a camelCase or PascalCase string to snake_case.
func ConvertToSnakeCase(input string) string {
	var matchChars = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAlpha = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchChars.ReplaceAllString(input, "${1}_${2}")
	snake = matchAlpha.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// Converts a snake_case string to camelCase.
func ConvertSnakeToCamel(snake string) string {
	re := regexp.MustCompile("(_[a-z])")
	return re.ReplaceAllStringFunc(snake, func(s string) string {
		return strings.ToUpper(s[1:])
	})
}

// Converts a snake_case string to PascalCase.
func ConvertSnakeToPascal(snake string) string {
	c := ConvertSnakeToCamel(snake)
	p := strings.ToUpper(string(c[0])) + c[1:]

	return p
}

// Returns the first path found in the input line string as a pathlib.Path.
func FindPathInLine(inputLine string) pathlib.Path {
	pathPattern := regexp.MustCompile(`[A-Za-z]:[\\/](?:[\w.-]+[\\/]*)+`)
	found := pathPattern.FindString(inputLine)
	return *pathlib.NewPath(found)
}
