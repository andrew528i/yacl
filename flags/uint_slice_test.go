package flags

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_UintSlice(t *testing.T) {
	type ChildPorts struct {
		PostgresPorts []uint
	}

	type Ports struct {
		PostgresPorts   []uint32
		ClickhousePorts []uint64
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
				"-postgres-ports", "5432", "-postgres-ports", "5433", "-postgres-ports", "5434",
				"-clickhouse-ports", "8123", "-clickhouse-ports", "8124", "-clickhouse-ports", "8125",
				"-child-postgres-ports", "1234", "-child-postgres-ports", "1235",
			},
			expected: Ports{
				PostgresPorts:   []uint32{5432, 5433, 5434},
				ClickhousePorts: []uint64{8123, 8124, 8125},
				Child: ChildPorts{
					PostgresPorts: []uint{1234, 1235},
				},
			},
		}, {
			name: "no-postgres-ports",
			args: []string{
				"cmd",
				"-clickhouse-ports", "8123", "-clickhouse-ports", "8124", "-clickhouse-ports", "8125",
				"-child-postgres-ports", "1234", "-child-postgres-ports", "1235",
			},
			expected: Ports{
				ClickhousePorts: []uint64{8123, 8124, 8125},
				Child: ChildPorts{
					PostgresPorts: []uint{1234, 1235},
				},
			},
		}, {
			name: "no-child-ports",
			args: []string{
				"cmd",
				"-postgres-ports", "5432", "-postgres-ports", "5433", "-postgres-ports", "5434",
				"-clickhouse-ports", "8123", "-clickhouse-ports", "8124", "-clickhouse-ports", "8125",
			},
			expected: Ports{
				PostgresPorts:   []uint32{5432, 5433, 5434},
				ClickhousePorts: []uint64{8123, 8124, 8125},
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

			ports, err := Parse[Ports](DefaultParams())
			assert.NoError(t, err)

			assert.Equal(t, tc.expected, *ports)
		})
	}

	os.Args = originalArgs
}
