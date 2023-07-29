package flags

import (
	"strings"
	"unicode"
)

func camelCaseToSlice(s string) (result []string) {
	// Iterate through each character in the string
	start := 0
	for i := 1; i < len(s); i++ {
		// If the current character is uppercase, add the previous word to the slice
		if unicode.IsUpper(rune(s[i])) {
			result = append(result, s[start:i])
			start = i
		}
	}

	// Add the final word to the slice
	result = append(result, strings.ToLower(s[start:]))

	return
}
