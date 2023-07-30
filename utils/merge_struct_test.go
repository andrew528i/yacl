package utils

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeStruct(t *testing.T) {
	type Address struct {
		Street string
	}

	type Person struct {
		Name    string
		Age     int
		Address Address
		Tags    []string
	}

	tests := []struct {
		name         string
		target       Person
		source       Person
		expected     Person
		expectChange bool
	}{
		{
			name:         "no-change",
			target:       Person{Name: "John", Age: 30, Tags: []string{"hello", "world"}, Address: Address{Street: "123 Main St"}},
			source:       Person{Name: "", Age: 0, Address: Address{Street: ""}},
			expected:     Person{Name: "John", Age: 30, Tags: []string{"hello", "world"}, Address: Address{Street: "123 Main St"}},
			expectChange: false,
		},
		{
			name:         "no-age",
			target:       Person{Name: "John", Age: 30, Address: Address{Street: "123 Main St"}},
			source:       Person{Name: "Doe", Age: 0, Tags: []string{"hello", "world"}, Address: Address{Street: "456 Second St"}},
			expected:     Person{Name: "Doe", Age: 30, Tags: []string{"hello", "world"}, Address: Address{Street: "456 Second St"}},
			expectChange: true,
		},
		{
			name:         "complete",
			target:       Person{Name: "John", Age: 30, Address: Address{Street: "123 Main St"}},
			source:       Person{Name: "Doe", Age: 25, Address: Address{Street: "456 Second St"}},
			expected:     Person{Name: "Doe", Age: 25, Address: Address{Street: "456 Second St"}},
			expectChange: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of the target struct
			target := tt.target

			// Merge the source struct into the target struct
			MergeStruct(&target, &tt.source)

			// Check if the target struct matches the expected struct
			assert.Equal(t, tt.expected, target)

			// Check if the target struct was expected to change
			assert.Equal(t, tt.expectChange, !reflect.DeepEqual(target, tt.target))
		})
	}
}
