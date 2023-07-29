package utils

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalkStruct(t *testing.T) {
	type Address struct {
		Street  string
		City    string `yacl:"testing"`
		Country string
	}

	type Person struct {
		Name    string
		Age     int
		Address Address
	}

	person := Person{
		Name: "John Doe",
		Age:  30,
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
		},
	}

	// Create a slice to store the field paths and values
	var fieldPaths []string
	var fieldValues []reflect.Value
	var meetTag bool

	// Define the callback function to store the field paths and values
	callback := func(fieldPath []string, field reflect.Value, fieldTag *reflect.StructTag) {
		fieldPaths = append(fieldPaths, strings.Join(fieldPath, "."))
		fieldValues = append(fieldValues, field)
		if fieldTag.Get("yacl") == "testing" {
			meetTag = true
		}
	}

	// Call the WalkStruct function with the person struct and the callback function
	WalkStruct(&person, callback)

	assert.True(t, meetTag)

	// Perform assertions on the field paths and values
	expectedFieldPaths := []string{
		"Name",
		"Age",
		"Address.Street",
		"Address.City",
		"Address.Country",
	}
	expectedFieldValues := []reflect.Value{
		reflect.ValueOf("John Doe"),
		reflect.ValueOf(30),
		reflect.ValueOf("123 Main St"),
		reflect.ValueOf("New York"),
		reflect.ValueOf("USA"),
	}

	assert.Equal(t, expectedFieldPaths, fieldPaths, "Field paths should match")

	for i := 0; i < len(expectedFieldValues); i++ {
		assert.Equal(t, expectedFieldValues[i].String(), fieldValues[i].String())
	}
}
