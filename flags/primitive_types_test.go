package flags

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_String(t *testing.T) {
	type Address struct {
		Street string `yacl:"street-alias"`
	}

	type Person struct {
		Name    string
		Address Address
	}

	testCases := []struct {
		name     string
		args     []string
		expected Person
	}{
		{
			name: "complete",
			args: []string{
				"cmd",
				"-name", "John Doe",
				"-street-alias", "Main St.",
			},
			expected: Person{
				Name: "John Doe",
				Address: Address{
					Street: "Main St.",
				},
			},
		}, {
			name: "no-address-street",
			args: []string{
				"cmd",
				"-name", "John Doe",
			},
			expected: Person{Name: "John Doe"},
		}, {
			name: "no-name",
			args: []string{
				"cmd",
				"-street-alias", "Main St.",
			},
			expected: Person{
				Address: Address{
					Street: "Main St.",
				},
			},
		}, {
			name:     "empty",
			args:     []string{"cmd"},
			expected: Person{},
		},
	}

	originalArgs := os.Args

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			os.Args = tc.args

			person, err := Parse[Person](DefaultParams())
			assert.NoError(t, err)

			assert.Equal(t, tc.expected, *person)
		})
	}

	os.Args = originalArgs
}

func TestParse_Bool(t *testing.T) {
	type Child struct {
		B bool
	}

	type Parent struct {
		A bool
		C bool

		Child Child
	}

	testCases := []struct {
		name     string
		args     []string
		expected Parent
	}{
		{
			name: "complete",
			args: []string{"cmd", "-c", "-a", "-child-b"},
			expected: Parent{
				A: true,
				C: true,
				Child: Child{
					B: true,
				},
			},
		}, {
			name: "no-parent",
			args: []string{"cmd", "-child-b"},
			expected: Parent{
				Child: Child{
					B: true,
				},
			},
		}, {
			name:     "no-child",
			args:     []string{"cmd", "-a", "-c"},
			expected: Parent{A: true, C: true},
		}, {
			name:     "empty",
			args:     []string{"cmd"},
			expected: Parent{},
		},
	}

	originalArgs := os.Args

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			os.Args = tc.args

			parent, err := Parse[Parent](DefaultParams())
			assert.NoError(t, err)

			assert.Equal(t, tc.expected, *parent)
		})
	}

	os.Args = originalArgs
}

func TestParse_Uint(t *testing.T) {
	type PreviousGrade struct {
		Score uint64
	}

	type Grade struct {
		Score uint32

		Previous PreviousGrade
	}

	type Person struct {
		Age    uint32
		Weight uint64

		Grade Grade
	}

	testCases := []struct {
		name     string
		args     []string
		expected Person
	}{
		{
			name: "complete",
			args: []string{
				"cmd",
				"-age", "19",
				"-weight", "75",
				"-grade-score", "10",
				"-grade-previous-score", "15",
			},
			expected: Person{
				Age:    19,
				Weight: 75,
				Grade: Grade{
					Score: 10,
					Previous: PreviousGrade{
						Score: 15,
					},
				},
			},
		}, {
			name: "no-age",
			args: []string{
				"cmd",
				"-weight", "75",
				"-grade-score", "10",
				"-grade-previous-score", "15",
			},
			expected: Person{
				Weight: 75,
				Grade: Grade{
					Score: 10,
					Previous: PreviousGrade{
						Score: 15,
					},
				},
			},
		}, {
			name: "no-weight-grade",
			args: []string{
				"cmd",
				"-age", "19",
				"-grade-previous-score", "15",
			},
			expected: Person{
				Age: 19,
				Grade: Grade{
					Previous: PreviousGrade{
						Score: 15,
					},
				},
			},
		}, {
			name:     "empty",
			args:     []string{"cmd"},
			expected: Person{},
		},
	}

	originalArgs := os.Args

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			os.Args = tc.args

			person, err := Parse[Person](DefaultParams())
			assert.NoError(t, err)

			assert.Equal(t, tc.expected, *person)
		})
	}

	os.Args = originalArgs
}

func TestParse_Int(t *testing.T) {
	type PreviousGrade struct {
		Score int64
	}

	type Grade struct {
		Score int32

		Previous PreviousGrade
	}

	type Person struct {
		Age    int32
		Weight int64

		Grade Grade
	}

	testCases := []struct {
		name     string
		args     []string
		expected Person
	}{
		{
			name: "complete",
			args: []string{
				"cmd",
				"-age", "-19",
				"-weight", "-75",
				"-grade-score", "-10",
				"-grade-previous-score", "-15",
			},
			expected: Person{
				Age:    -19,
				Weight: -75,
				Grade: Grade{
					Score: -10,
					Previous: PreviousGrade{
						Score: -15,
					},
				},
			},
		}, {
			name: "no-age",
			args: []string{
				"cmd",
				"-weight", "-75",
				"-grade-score", "-10",
				"-grade-previous-score", "-15",
			},
			expected: Person{
				Weight: -75,
				Grade: Grade{
					Score: -10,
					Previous: PreviousGrade{
						Score: -15,
					},
				},
			},
		}, {
			name: "no-weight-grade",
			args: []string{
				"cmd",
				"-age", "-19",
				"-grade-previous-score", "-15",
			},
			expected: Person{
				Age: -19,
				Grade: Grade{
					Previous: PreviousGrade{
						Score: -15,
					},
				},
			},
		}, {
			name:     "empty",
			args:     []string{"cmd"},
			expected: Person{},
		},
	}

	originalArgs := os.Args

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			os.Args = tc.args

			person, err := Parse[Person](DefaultParams())
			assert.NoError(t, err)

			assert.Equal(t, tc.expected, *person)
		})
	}

	os.Args = originalArgs
}

func TestParse_Float64(t *testing.T) {
	type YesterdayTemperature struct {
		Value float64
	}

	type Temperature struct {
		Value     float64
		Yesterday YesterdayTemperature
	}

	testCases := []struct {
		name     string
		args     []string
		expected Temperature
	}{
		{
			name: "complete",
			args: []string{
				"cmd",
				"-value", "23.75",
				"-yesterday-value", "21.13",
			},
			expected: Temperature{
				Value: 23.75,
				Yesterday: YesterdayTemperature{
					Value: 21.13,
				},
			},
		}, {
			name: "no-value",
			args: []string{
				"cmd",
				"-yesterday-value", "21.13",
			},
			expected: Temperature{
				Yesterday: YesterdayTemperature{
					Value: 21.13,
				},
			},
		}, {
			name:     "no-yesterday",
			args:     []string{"cmd", "-value", "23.75"},
			expected: Temperature{Value: 23.75},
		}, {
			name:     "empty",
			args:     []string{"cmd"},
			expected: Temperature{},
		},
	}

	originalArgs := os.Args

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			os.Args = tc.args

			temperature, err := Parse[Temperature](DefaultParams())
			assert.NoError(t, err)

			assert.Equal(t, tc.expected, *temperature)
		})
	}

	os.Args = originalArgs
}
