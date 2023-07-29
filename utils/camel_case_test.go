package utils

import (
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

func TestCamelCaseToSlice(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "single word",
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			name:     "multiple words",
			input:    "helloWorld",
			expected: []string{"hello", "world"},
		},
		{
			name:     "words separated by digit",
			input:    "hello123World",
			expected: []string{"hello", "123", "world"},
		},
		{
			name:     "words separated by special character",
			input:    "hello$world",
			expected: []string{"hello", "$", "world"},
		},
		{
			name:     "words separated by emoji",
			input:    "hello😀world",
			expected: []string{"hello", "😀", "world"},
		},
		{
			name:     "non-ascii characters",
			input:    "こんにちは世界",
			expected: []string{"こんにちは世界"},
		},
		{
			name:     "invalid utf8 string",
			input:    "hello\x80world",
			expected: []string{"hello\x80world"},
		},
		{
			name:     "acronym",
			input:    "HTTPRequest",
			expected: []string{"http", "request"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := CamelCaseToSlice(tc.input)
			assert.Equal(t, tc.expected, actual)

			// Additional assertions for valid utf8 strings
			if utf8.ValidString(tc.input) {
				for _, entry := range actual {
					// Assert that each entry is lowercase
					assert.True(t, strings.ToLower(entry) == entry)

					// Assert that each entry does not contain any uppercase letters
					for _, r := range entry {
						assert.False(t, unicode.IsUpper(r))
					}
				}
			}
		})
	}
}
