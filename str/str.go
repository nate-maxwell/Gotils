package str

import (
	"regexp"
	"strings"
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
func ConvertSnakeToPascale(snake string) string {
	c := ConvertSnakeToCamel(snake)
	p := strings.ToUpper(string(c[0])) + c[1:]

	return p
}
