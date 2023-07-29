package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseString(t *testing.T) {
	type DatabaseConfig struct {
		Hostname string
		Username string
		Password string
	}

	type Config struct {
		HTTPBindAddress string

		Database DatabaseConfig
	}

	testCases := []struct {
		name     string
		vars     map[string]string
		expected Config
	}{
		{
			name: "complete",
			vars: map[string]string{
				"HTTP_BIND_ADDRESS": "localhost:8080",
				"DATABASE_HOSTNAME": "127.0.0.1",
				"DATABASE_USERNAME": "test-username",
				"DATABASE_PASSWORD": "test-password",
			},
			expected: Config{
				HTTPBindAddress: "localhost:8080",
				Database: DatabaseConfig{
					Hostname: "127.0.0.1",
					Username: "test-username",
					Password: "test-password",
				},
			},
		}, {
			name: "no-http-bind-address",
			vars: map[string]string{
				"DATABASE_HOSTNAME": "127.0.0.1",
				"DATABASE_USERNAME": "test-username",
				"DATABASE_PASSWORD": "test-password",
			},
			expected: Config{
				Database: DatabaseConfig{
					Hostname: "127.0.0.1",
					Username: "test-username",
					Password: "test-password",
				},
			},
		}, {
			name: "no-database-hostname",
			vars: map[string]string{
				"HTTP_BIND_ADDRESS": "localhost:8080",
				"DATABASE_USERNAME": "test-username",
				"DATABASE_PASSWORD": "test-password",
			},
			expected: Config{
				HTTPBindAddress: "localhost:8080",
				Database: DatabaseConfig{
					Username: "test-username",
					Password: "test-password",
				},
			},
		}, {
			name:     "empty",
			vars:     map[string]string{},
			expected: Config{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.vars {
				assert.NoError(t, os.Setenv(k, v))
			}

			cfg, err := Parse[Config](DefaultParams())
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, *cfg)

			os.Clearenv()
		})
	}
}

func TestParse_Bool(t *testing.T) {
	type TerminalConfig struct {
		BellSound bool
	}

	type Config struct {
		ShouldPanic bool
		CanStop     bool
		Terminal    TerminalConfig
	}

	testCases := []struct {
		name     string
		vars     map[string]string
		expected Config
	}{
		{
			name: "complete",
			vars: map[string]string{
				"SHOULD_PANIC":        "true",
				"CAN_STOP":            "true",
				"TERMINAL_BELL_SOUND": "true",
			},
			expected: Config{
				ShouldPanic: true,
				CanStop:     true,
				Terminal: TerminalConfig{
					BellSound: true,
				},
			},
		}, {
			name: "no-can-stop",
			vars: map[string]string{
				"SHOULD_PANIC":        "true",
				"TERMINAL_BELL_SOUND": "true",
			},
			expected: Config{
				ShouldPanic: true,
				Terminal: TerminalConfig{
					BellSound: true,
				},
			},
		}, {
			name: "no-bell-sound",
			vars: map[string]string{
				"SHOULD_PANIC": "true",
				"CAN_STOP":     "true",
			},
			expected: Config{
				ShouldPanic: true,
				CanStop:     true,
				Terminal:    TerminalConfig{},
			},
		}, {
			name:     "empty",
			vars:     map[string]string{},
			expected: Config{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.vars {
				assert.NoError(t, os.Setenv(k, v))
			}

			cfg, err := Parse[Config](DefaultParams())
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, *cfg)

			os.Clearenv()
		})
	}
}

func TestParse_Uint(t *testing.T) {
	type DatabaseConfig struct {
		Port    uint64
		Timeout uint32
	}

	type Config struct {
		HTTPPort   uint
		Postgres   DatabaseConfig
		Clickhouse DatabaseConfig
	}

	testCases := []struct {
		name     string
		vars     map[string]string
		expected Config
	}{
		{
			name: "complete",
			vars: map[string]string{
				"HTTP_PORT":          "8081",
				"POSTGRES_PORT":      "5432",
				"POSTGRES_TIMEOUT":   "10000",
				"CLICKHOUSE_PORT":    "8123",
				"CLICKHOUSE_TIMEOUT": "15000",
			},
			expected: Config{
				HTTPPort: 8081,
				Postgres: DatabaseConfig{
					Port:    5432,
					Timeout: 10000,
				},
				Clickhouse: DatabaseConfig{
					Port:    8123,
					Timeout: 15000,
				},
			},
		}, {
			name: "no-http-port",
			vars: map[string]string{
				"POSTGRES_PORT":      "5432",
				"POSTGRES_TIMEOUT":   "10000",
				"CLICKHOUSE_PORT":    "8123",
				"CLICKHOUSE_TIMEOUT": "15000",
			},
			expected: Config{
				Postgres: DatabaseConfig{
					Port:    5432,
					Timeout: 10000,
				},
				Clickhouse: DatabaseConfig{
					Port:    8123,
					Timeout: 15000,
				},
			},
		}, {
			name: "no-postgres-port",
			vars: map[string]string{
				"HTTP_PORT":          "8081",
				"POSTGRES_TIMEOUT":   "10000",
				"CLICKHOUSE_PORT":    "8123",
				"CLICKHOUSE_TIMEOUT": "15000",
			},
			expected: Config{
				HTTPPort: 8081,
				Postgres: DatabaseConfig{
					Timeout: 10000,
				},
				Clickhouse: DatabaseConfig{
					Port:    8123,
					Timeout: 15000,
				},
			},
		}, {
			name:     "empty",
			vars:     map[string]string{},
			expected: Config{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.vars {
				assert.NoError(t, os.Setenv(k, v))
			}

			cfg, err := Parse[Config](DefaultParams())
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, *cfg)

			os.Clearenv()
		})
	}
}

func TestParse_Int(t *testing.T) {
	type DatabaseConfig struct {
		Port    int64
		Timeout int32
	}

	type Config struct {
		HTTPPort   int
		Postgres   DatabaseConfig
		Clickhouse DatabaseConfig
	}

	testCases := []struct {
		name     string
		vars     map[string]string
		expected Config
	}{
		{
			name: "complete",
			vars: map[string]string{
				"HTTP_PORT":          "-8081",
				"POSTGRES_PORT":      "-5432",
				"POSTGRES_TIMEOUT":   "-10000",
				"CLICKHOUSE_PORT":    "-8123",
				"CLICKHOUSE_TIMEOUT": "-15000",
			},
			expected: Config{
				HTTPPort: -8081,
				Postgres: DatabaseConfig{
					Port:    -5432,
					Timeout: -10000,
				},
				Clickhouse: DatabaseConfig{
					Port:    -8123,
					Timeout: -15000,
				},
			},
		}, {
			name: "no-http-port",
			vars: map[string]string{
				"POSTGRES_PORT":      "-5432",
				"POSTGRES_TIMEOUT":   "-10000",
				"CLICKHOUSE_PORT":    "-8123",
				"CLICKHOUSE_TIMEOUT": "-15000",
			},
			expected: Config{
				Postgres: DatabaseConfig{
					Port:    -5432,
					Timeout: -10000,
				},
				Clickhouse: DatabaseConfig{
					Port:    -8123,
					Timeout: -15000,
				},
			},
		}, {
			name: "no-postgres-port",
			vars: map[string]string{
				"HTTP_PORT":          "-8081",
				"POSTGRES_TIMEOUT":   "-10000",
				"CLICKHOUSE_PORT":    "-8123",
				"CLICKHOUSE_TIMEOUT": "-15000",
			},
			expected: Config{
				HTTPPort: -8081,
				Postgres: DatabaseConfig{
					Timeout: -10000,
				},
				Clickhouse: DatabaseConfig{
					Port:    -8123,
					Timeout: -15000,
				},
			},
		}, {
			name:     "empty",
			vars:     map[string]string{},
			expected: Config{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.vars {
				assert.NoError(t, os.Setenv(k, v))
			}

			cfg, err := Parse[Config](DefaultParams())
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, *cfg)

			os.Clearenv()
		})
	}
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
		vars     map[string]string
		expected Temperature
	}{
		{
			name: "complete",
			vars: map[string]string{
				"VALUE":           "21.7777",
				"YESTERDAY_VALUE": "93.1123333",
			},
			expected: Temperature{
				Value: 21.7777,
				Yesterday: YesterdayTemperature{
					Value: 93.1123333,
				},
			},
		}, {
			name: "no-value",
			vars: map[string]string{
				"YESTERDAY_VALUE": "93.1123333",
			},
			expected: Temperature{
				Yesterday: YesterdayTemperature{
					Value: 93.1123333,
				},
			},
		}, {
			name: "no-yesterday-value",
			vars: map[string]string{
				"VALUE": "21.7777",
			},
			expected: Temperature{
				Value:     21.7777,
				Yesterday: YesterdayTemperature{},
			},
		}, {
			name:     "empty",
			vars:     map[string]string{},
			expected: Temperature{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.vars {
				assert.NoError(t, os.Setenv(k, v))
			}

			cfg, err := Parse[Temperature](DefaultParams())
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, *cfg)

			os.Clearenv()
		})
	}
}

func TestParse_StringSlice(t *testing.T) {
	type ChildConfig struct {
		Tags []string
	}

	type Config struct {
		Tags  []string
		Child ChildConfig
	}

	testCases := []struct {
		name     string
		vars     map[string]string
		expected Config
	}{
		{
			name: "complete",
			vars: map[string]string{
				"TAGS":       "quick,brown,fox",
				"CHILD_TAGS": "hello,world",
			},
			expected: Config{
				Tags: []string{"quick", "brown", "fox"},
				Child: ChildConfig{
					Tags: []string{"hello", "world"},
				},
			},
		}, {
			name: "no-tags",
			vars: map[string]string{
				"CHILD_TAGS": "hello,world",
			},
			expected: Config{
				Child: ChildConfig{
					Tags: []string{"hello", "world"},
				},
			},
		}, {
			name: "no-child-tags",
			vars: map[string]string{
				"TAGS": "quick,brown,fox",
			},
			expected: Config{
				Tags:  []string{"quick", "brown", "fox"},
				Child: ChildConfig{},
			},
		}, {
			name:     "empty",
			vars:     map[string]string{},
			expected: Config{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.vars {
				assert.NoError(t, os.Setenv(k, v))
			}

			cfg, err := Parse[Config](DefaultParams())
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, *cfg)

			os.Clearenv()
		})
	}
}

func TestParse_BoolSlice(t *testing.T) {
	type ChildConfig struct {
		Signals []bool
	}

	type Config struct {
		Signals []bool
		Child   ChildConfig
	}

	testCases := []struct {
		name     string
		vars     map[string]string
		expected Config
	}{
		{
			name: "complete",
			vars: map[string]string{
				"SIGNALS":       "true,true,false,true",
				"CHILD_SIGNALS": "false,true,true,false",
			},
			expected: Config{
				Signals: []bool{true, true, false, true},
				Child: ChildConfig{
					Signals: []bool{false, true, true, false},
				},
			},
		}, {
			name: "no-signals",
			vars: map[string]string{
				"CHILD_SIGNALS": "false,true,true,false",
			},
			expected: Config{
				Child: ChildConfig{
					Signals: []bool{false, true, true, false},
				},
			},
		}, {
			name: "no-child-signals",
			vars: map[string]string{
				"SIGNALS": "true,true,false,true",
			},
			expected: Config{
				Signals: []bool{true, true, false, true},
				Child:   ChildConfig{},
			},
		}, {
			name:     "empty",
			vars:     map[string]string{},
			expected: Config{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.vars {
				assert.NoError(t, os.Setenv(k, v))
			}

			cfg, err := Parse[Config](DefaultParams())
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, *cfg)

			os.Clearenv()
		})
	}
}

func TestParse_UintSlice(t *testing.T) {
	type DatabaseConfig struct {
		Ports []uint32
	}

	type Config struct {
		Postgres        DatabaseConfig
		ClickhousePorts []uint64
		Ports           []uint
	}

	testCases := []struct {
		name     string
		vars     map[string]string
		expected Config
	}{
		{
			name: "complete",
			vars: map[string]string{
				"PORTS":            "1234,1235,1236,1237",
				"CLICKHOUSE_PORTS": "8123,8124,8125",
				"POSTGRES_PORTS":   "5432,5433,5434,5435",
			},
			expected: Config{
				Postgres: DatabaseConfig{
					Ports: []uint32{5432, 5433, 5434, 5435},
				},
				ClickhousePorts: []uint64{8123, 8124, 8125},
				Ports:           []uint{1234, 1235, 1236, 1237},
			},
		}, {
			name: "no-ports",
			vars: map[string]string{
				"CLICKHOUSE_PORTS": "8123,8124,8125",
				"POSTGRES_PORTS":   "5432,5433,5434,5435",
			},
			expected: Config{
				Postgres: DatabaseConfig{
					Ports: []uint32{5432, 5433, 5434, 5435},
				},
				ClickhousePorts: []uint64{8123, 8124, 8125},
			},
		}, {
			name: "no-postgres-ports",
			vars: map[string]string{
				"PORTS":            "1234,1235,1236,1237",
				"CLICKHOUSE_PORTS": "8123,8124,8125",
			},
			expected: Config{
				Postgres:        DatabaseConfig{},
				ClickhousePorts: []uint64{8123, 8124, 8125},
				Ports:           []uint{1234, 1235, 1236, 1237},
			},
		}, {
			name:     "empty",
			vars:     map[string]string{},
			expected: Config{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.vars {
				assert.NoError(t, os.Setenv(k, v))
			}

			cfg, err := Parse[Config](DefaultParams())
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, *cfg)

			os.Clearenv()
		})
	}
}

func TestParse_IntSlice(t *testing.T) {
	type DatabaseConfig struct {
		Ports []int32
	}

	type Config struct {
		Postgres        DatabaseConfig
		ClickhousePorts []int64
		Ports           []int
	}

	testCases := []struct {
		name     string
		vars     map[string]string
		expected Config
	}{
		{
			name: "complete",
			vars: map[string]string{
				"PORTS":            "-1234,-1235,-1236,-1237",
				"CLICKHOUSE_PORTS": "-8123,-8124,-8125",
				"POSTGRES_PORTS":   "-5432,-5433,-5434,-5435",
			},
			expected: Config{
				Postgres: DatabaseConfig{
					Ports: []int32{-5432, -5433, -5434, -5435},
				},
				ClickhousePorts: []int64{-8123, -8124, -8125},
				Ports:           []int{-1234, -1235, -1236, -1237},
			},
		}, {
			name: "no-ports",
			vars: map[string]string{
				"CLICKHOUSE_PORTS": "-8123,-8124,-8125",
				"POSTGRES_PORTS":   "-5432,-5433,-5434,-5435",
			},
			expected: Config{
				Postgres: DatabaseConfig{
					Ports: []int32{-5432, -5433, -5434, -5435},
				},
				ClickhousePorts: []int64{-8123, -8124, -8125},
			},
		}, {
			name: "no-postgres-ports",
			vars: map[string]string{
				"PORTS":            "-1234,-1235,-1236,-1237",
				"CLICKHOUSE_PORTS": "-8123,-8124,-8125",
			},
			expected: Config{
				Postgres:        DatabaseConfig{},
				ClickhousePorts: []int64{-8123, -8124, -8125},
				Ports:           []int{-1234, -1235, -1236, -1237},
			},
		}, {
			name:     "empty",
			vars:     map[string]string{},
			expected: Config{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.vars {
				assert.NoError(t, os.Setenv(k, v))
			}

			cfg, err := Parse[Config](DefaultParams())
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, *cfg)

			os.Clearenv()
		})
	}
}

func TestParse_Float64Slice(t *testing.T) {
	type ChildConfig struct {
		Values []float64
	}

	type Config struct {
		Values []float64
		Child  ChildConfig
	}

	testCases := []struct {
		name     string
		vars     map[string]string
		expected Config
	}{
		{
			name: "complete",
			vars: map[string]string{
				"VALUES":       "21.7733,-2299.122311,-1.0000019",
				"CHILD_VALUES": "222.2233129999900001,-19.229939",
			},
			expected: Config{
				Values: []float64{21.7733, -2299.122311, -1.0000019},
				Child: ChildConfig{
					Values: []float64{222.2233129999900001, -19.229939},
				},
			},
		}, {
			name: "no-values",
			vars: map[string]string{
				"CHILD_VALUES": "222.2233129999900001,-19.229939",
			},
			expected: Config{
				Child: ChildConfig{
					Values: []float64{222.2233129999900001, -19.229939},
				},
			},
		}, {
			name: "complete",
			vars: map[string]string{
				"VALUES": "21.7733,-2299.122311,-1.0000019",
			},
			expected: Config{
				Values: []float64{21.7733, -2299.122311, -1.0000019},
				Child:  ChildConfig{},
			},
		}, {
			name:     "empty",
			vars:     map[string]string{},
			expected: Config{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.vars {
				assert.NoError(t, os.Setenv(k, v))
			}

			cfg, err := Parse[Config](DefaultParams())
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, *cfg)

			os.Clearenv()
		})
	}
}
