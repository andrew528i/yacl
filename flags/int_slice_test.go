package flags

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromFlags_IntSlice(t *testing.T) {
	type ChildPorts struct {
		PostgresPorts []int
	}

	type Ports struct {
		PostgresPorts   []int32
		ClickhousePorts []int64
		Child           ChildPorts
	}

	testCases := []struct {
		name     string
		args     []string
		expected Ports
	}{
		{
			name: "complete",
			args: []string{
				"cmd",
				"-postgres-ports", "-5432", "-postgres-ports", "-5433", "-postgres-ports", "-5434",
				"-clickhouse-ports", "-8123", "-clickhouse-ports", "-8124", "-clickhouse-ports", "-8125",
				"-child-postgres-ports", "-1234", "-child-postgres-ports", "-1235",
			},
			expected: Ports{
				PostgresPorts:   []int32{-5432, -5433, -5434},
				ClickhousePorts: []int64{-8123, -8124, -8125},
				Child: ChildPorts{
					PostgresPorts: []int{-1234, -1235},
				},
			},
		}, {
			name: "no-clickhouse-ports",
			args: []string{
				"cmd",
				"-postgres-ports", "-5432", "-postgres-ports", "-5433", "-postgres-ports", "-5434",
				"-child-postgres-ports", "-1234", "-child-postgres-ports", "-1235",
			},
			expected: Ports{
				PostgresPorts: []int32{-5432, -5433, -5434},
				Child: ChildPorts{
					PostgresPorts: []int{-1234, -1235},
				},
			},
		}, {
			name: "no-child-ports",
			args: []string{
				"cmd",
				"-postgres-ports", "-5432", "-postgres-ports", "-5433", "-postgres-ports", "-5434",
				"-clickhouse-ports", "-8123", "-clickhouse-ports", "-8124", "-clickhouse-ports", "-8125",
			},
			expected: Ports{
				PostgresPorts:   []int32{-5432, -5433, -5434},
				ClickhousePorts: []int64{-8123, -8124, -8125},
			},
		}, {
			name:     "empty",
			args:     []string{"cmd"},
			expected: Ports{},
		},
	}

	originalArgs := os.Args

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			os.Args = tc.args

			ports := FromFlags[Ports](DefaultFlagParams())

			assert.Equal(t, tc.expected, *ports)
		})
	}

	os.Args = originalArgs
}
