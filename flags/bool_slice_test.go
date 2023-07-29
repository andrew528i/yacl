package flags

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromFlags_BoolSlice(t *testing.T) {
	type MyStructChild struct {
		Marks []bool
	}

	type MyStruct struct {
		Marks []bool
		Child MyStructChild
	}

	testCases := []struct {
		name     string
		args     []string
		expected MyStruct
	}{
		{
			name: "complete",
			args: []string{
				"cmd",
				"-marks", "true", "-marks", "false", "-marks", "true",
				"-child-marks", "false", "-child-marks", "true", "-child-marks", "false",
			},
			expected: MyStruct{
				Marks: []bool{true, false, true},
				Child: MyStructChild{
					Marks: []bool{false, true, false},
				},
			},
		}, {
			name: "no-marks",
			args: []string{
				"cmd",
				"-child-marks", "false", "-child-marks", "true", "-child-marks", "false",
			},
			expected: MyStruct{
				Child: MyStructChild{
					Marks: []bool{false, true, false},
				},
			},
		}, {
			name: "no-child-marks",
			args: []string{
				"cmd",
				"-marks", "true", "-marks", "false", "-marks", "true",
			},
			expected: MyStruct{
				Marks: []bool{true, false, true},
				Child: MyStructChild{},
			},
		}, {
			name:     "empty",
			args:     []string{"cmd"},
			expected: MyStruct{},
		},
	}

	originalArgs := os.Args

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			os.Args = tc.args

			myStruct := FromFlags[MyStruct](DefaultFlagParams())

			assert.Equal(t, tc.expected, *myStruct)
		})
	}

	os.Args = originalArgs
}
