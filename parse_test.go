package yacl

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/andrew528i/yacl/file"
	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack/v5"
	"gopkg.in/yaml.v3"
)

func TestParse(t *testing.T) {
	type Database struct {
		Hostname string
	}

	type Config struct {
		HTTPPort uint
		Postgres Database
	}

	// Prepare default config
	defaultConfigs := []*Config{
		{
			HTTPPort: 8086,
			Postgres: Database{
				Hostname: "localhost-default",
			},
		},
	}

	// Prepare yaml file
	yamlCfg := &Config{
		HTTPPort: 8083,
		Postgres: Database{
			Hostname: "localhost-yaml",
		},
	}

	yamlBytes, err := yaml.Marshal(yamlCfg)
	assert.NoError(t, err)

	// Prepare json file
	jsonCfg := &Config{
		HTTPPort: 8084,
		Postgres: Database{
			Hostname: "localhost-json",
		},
	}

	jsonBytes, err := json.Marshal(jsonCfg)
	assert.NoError(t, err)

	// Prepare binary file
	binaryCfg := &Config{
		HTTPPort: 8085,
		Postgres: Database{
			Hostname: "localhost-binary",
		},
	}

	binaryBytes, err := msgpack.Marshal(binaryCfg)
	assert.NoError(t, err)

	testCases := []struct {
		name           string
		args           []string
		vars           map[string]string
		yaml           []byte
		json           []byte
		binary         []byte
		defaultConfigs []*Config
		expected       Config
	}{
		{
			name: "complete",
			args: []string{
				"cmd",
				"-postgres-hostname", "localhost-flag",
				"-http-port", "8081",
			},
			vars: map[string]string{
				"POSTGRES_HOSTNAME": "localhost-env",
				"HTTP_PORT":         "8082",
			},
			yaml:           yamlBytes,
			json:           jsonBytes,
			binary:         binaryBytes,
			defaultConfigs: defaultConfigs,
			expected: Config{
				HTTPPort: 8081,
				Postgres: Database{
					Hostname: "localhost-flag",
				},
			},
		}, {
			name: "no-cli",
			args: []string{"cmd"},
			vars: map[string]string{
				"POSTGRES_HOSTNAME": "localhost-env",
				"HTTP_PORT":         "8082",
			},
			yaml:           yamlBytes,
			json:           jsonBytes,
			binary:         binaryBytes,
			defaultConfigs: defaultConfigs,
			expected: Config{
				HTTPPort: 8082,
				Postgres: Database{
					Hostname: "localhost-env",
				},
			},
		}, {
			name:           "no-env",
			args:           []string{"cmd"},
			vars:           map[string]string{},
			yaml:           yamlBytes,
			json:           jsonBytes,
			binary:         binaryBytes,
			defaultConfigs: defaultConfigs,
			expected: Config{
				HTTPPort: 8085,
				Postgres: Database{
					Hostname: "localhost-binary",
				},
			},
		}, {
			name:           "no-binary",
			args:           []string{"cmd"},
			vars:           map[string]string{},
			yaml:           yamlBytes,
			json:           jsonBytes,
			defaultConfigs: defaultConfigs,
			expected: Config{
				HTTPPort: 8084,
				Postgres: Database{
					Hostname: "localhost-json",
				},
			},
		}, {
			name:           "no-json",
			args:           []string{"cmd"},
			vars:           map[string]string{},
			yaml:           yamlBytes,
			defaultConfigs: defaultConfigs,
			expected: Config{
				HTTPPort: 8083,
				Postgres: Database{
					Hostname: "localhost-yaml",
				},
			},
		}, {
			name:           "no-yaml",
			args:           []string{"cmd"},
			vars:           map[string]string{},
			defaultConfigs: defaultConfigs,
			expected: Config{
				HTTPPort: 8086,
				Postgres: Database{
					Hostname: "localhost-default",
				},
			},
		}, {
			name:     "no-default",
			args:     []string{"cmd"},
			vars:     map[string]string{},
			expected: Config{},
		},
	}

	originalArgs := os.Args

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tempDir := t.TempDir()
			file.DefaultPath = tempDir

			// Set args
			if tc.args != nil {
				os.Args = tc.args
			}

			// Set env vars
			for k, v := range tc.vars {
				assert.NoError(t, os.Setenv(k, v))
			}

			// Write yaml file
			if tc.yaml != nil {
				assert.NoError(t, os.WriteFile(filepath.Join(tempDir, "config.yaml"), tc.yaml, 0644))
			}

			// Write json file
			if tc.json != nil {
				assert.NoError(t, os.WriteFile(filepath.Join(tempDir, "config.json"), tc.json, 0644))
			}

			// Write bin file
			if tc.binary != nil {
				assert.NoError(t, os.WriteFile(filepath.Join(tempDir, "config.bin"), tc.binary, 0644))
			}

			cfg, err := Parse(tc.defaultConfigs...)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, *cfg)

			// Cleanup
			if tc.args != nil {
				os.Args = originalArgs
			}

			os.Clearenv()
			file.DefaultPath = ""
		})
	}
}
