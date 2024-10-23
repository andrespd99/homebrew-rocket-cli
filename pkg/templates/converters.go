package templates

import (
	"strings"
	"unicode"
)

// TODO: Improve string converers...

// Function to convert a string to snake_case
func toSnakeCase(str string) string {
	var snake string
	for _, r := range str {
		if unicode.IsSpace(r) {
			snake += "_"
		} else {
			snake += string(unicode.ToLower(r))
		}
	}
	return snake
}

// Function to convert a string to camelCase
func toCamelCase(str string) string {
	parts := strings.Split(str, "_")
	for i := range parts {
		if i > 0 {
			parts[i] = strings.Title(parts[i])
		}
	}
	return strings.Join(parts, "")
}
