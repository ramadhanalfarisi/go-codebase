package helpers

import (
	"regexp"
	"strings"
)

func ToCamelCase(text string) string {
	// Replace non-alphanumeric characters with space
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	cleaned := reg.ReplaceAllString(text, " ")

	words := strings.Fields(cleaned)
	if len(words) == 0 {
		return ""
	}

	var result strings.Builder

	for i, word := range words {
		word = strings.ToLower(word)

		if i == 0 {
			result.WriteString(word)
		} else {
			result.WriteString(strings.ToUpper(word[:1]))
			if len(word) > 1 {
				result.WriteString(word[1:])
			}
		}
	}

	return result.String()
}