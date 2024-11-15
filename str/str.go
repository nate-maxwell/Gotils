package str

import (
	"regexp"
	"strings"
)

func ConvertToSnakeCase(input string) string {
	var matchChars = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAlpha = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchChars.ReplaceAllString(input, "${1}_${2}")
	snake = matchAlpha.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func ConvertSnakeToCamel(snake string) string {
	re := regexp.MustCompile("(_[a-z])")
	return re.ReplaceAllStringFunc(snake, func(s string) string {
		return strings.ToUpper(s[1:])
	})
}

func ConvertSnakeToPascale(snake string) string {
	re := regexp.MustCompile("(_[a-z])")
	result := re.ReplaceAllStringFunc(snake, func(s string) string {
		return strings.ToUpper(s[1:])
	})
	t := strings.ToUpper(string(result[0])) + result[1:]

	return t
}
