package blueprints

import (
	"strings"
	"unicode"
)

// TODO: Improve string converers...

// Function to convert a string to snake_case
func toSnakeCase(str string) string {
	var snake string
	for _, r := range str {
		if unicode.IsSpace(r) || unicode.Is(unicode.Hyphen, r) {
			snake += "_"
		} else {
			snake += string(unicode.ToLower(r))
		}
	}
	return snake
}

// Function to convert a string to camelCase
func toCamelCase(str string) string {
	str = strings.ToLower(str)
	r := strings.NewReplacer("-", " ", "_", " ")
	str = r.Replace(str)

	parts := strings.Split(str, " ")
	for i := range parts {
		if i > 0 {
			parts[i] = strings.Title(parts[i])
		}
	}

	return strings.Join(parts, "")
}

func toPascalCase(str string) string {
	r := strings.NewReplacer("-", " ", "_", " ")
	str = r.Replace(str)

	parts := strings.Split(str, " ")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}

func toTitleCase(str string) string {
	r := strings.NewReplacer("-", " ", "_", " ")
	str = r.Replace(str)

	return strings.Title(str)
}
