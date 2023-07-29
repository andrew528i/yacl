package flags

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromFlags_StringSlice(t *testing.T) {
	type ChildBlogPost struct {
		Tags []string
	}

	type BlogPost struct {
		Text  string
		Tags  []string
		Child ChildBlogPost
	}

	testCases := []struct {
		name     string
		args     []string
		expected BlogPost
	}{
		{
			name: "complete",
			args: []string{
				"cmd",
				"-text", "hello world",
				"-tags", "one", "-tags", "two", "-tags", "three",
				"-child-tags", "quick", "-child-tags", "brown", "-child-tags", "fox",
			},
			expected: BlogPost{
				Text: "hello world",
				Tags: []string{"one", "two", "three"},
				Child: ChildBlogPost{
					Tags: []string{"quick", "brown", "fox"},
				},
			},
		}, {
			name: "no-tags",
			args: []string{
				"cmd",
				"-text", "hello world",
				"-child-tags", "quick", "-child-tags", "brown", "-child-tags", "fox",
			},
			expected: BlogPost{
				Text: "hello world",
				Child: ChildBlogPost{
					Tags: []string{"quick", "brown", "fox"},
				},
			},
		}, {
			name: "no-child-tags",
			args: []string{
				"cmd",
				"-text", "hello world",
				"-tags", "one", "-tags", "two", "-tags", "three",
			},
			expected: BlogPost{
				Text:  "hello world",
				Tags:  []string{"one", "two", "three"},
				Child: ChildBlogPost{},
			},
		}, {
			name:     "empty",
			args:     []string{"cmd"},
			expected: BlogPost{},
		},
	}

	originalArgs := os.Args

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			os.Args = tc.args

			blogPost := FromFlags[BlogPost](DefaultFlagParams())

			assert.Equal(t, tc.expected, *blogPost)
		})
	}

	os.Args = originalArgs
}
