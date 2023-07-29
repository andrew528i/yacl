package flags

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_Float64Slice(t *testing.T) {
	type ChildTemperature struct {
		Values []float64
	}

	type Temperature struct {
		Values []float64
		Child  ChildTemperature
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
				"-values", "1.26", "-values", "1.49", "-values", "1.98",
				"-child-values", "21.23", "-child-values", "22.29", "-child-values", "25.91",
			},
			expected: Temperature{
				Values: []float64{1.26, 1.49, 1.98},
				Child: ChildTemperature{
					Values: []float64{21.23, 22.29, 25.91},
				},
			},
		}, {
			name: "no-values",
			args: []string{
				"cmd",
				"-child-values", "21.23123123", "-child-values", "22.29", "-child-values", "25.91",
			},
			expected: Temperature{
				Child: ChildTemperature{
					Values: []float64{21.23123123, 22.29, 25.91},
				},
			},
		}, {
			name: "no-child-values",
			args: []string{
				"cmd",
				"-values", "1.26321321321321321", "-values", "1.49", "-values", "1.98",
			},
			expected: Temperature{
				Values: []float64{1.26321321321321321, 1.49, 1.98},
			},
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
