package file

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack/v5"
)

func TestParseBinary(t *testing.T) {
	type DatabaseConfig struct {
		Port    uint64 `yaml:"port"`
		Restart bool   `yaml:"restart"`
	}

	type LoggerConfig struct {
		Levels []uint `yaml:"levels"`
	}

	type Config struct {
		Username        string         `yaml:"username"`
		MaxTemperatures []float64      `yaml:"max_temperatures"`
		Database        DatabaseConfig `yaml:"database"`
		Logger          LoggerConfig   `yaml:"logger"`
	}

	tempDir := t.TempDir()

	cfg := &Config{
		Username:        "some-testing",
		MaxTemperatures: []float64{12.3322222, 994.23122222222, 23.8172},
		Database: DatabaseConfig{
			Port:    5432,
			Restart: true,
		},
		Logger: LoggerConfig{
			Levels: []uint{3, 2, 1},
		},
	}

	tempFile := filepath.Join(tempDir, "config.bin")
	cfgData, err := msgpack.Marshal(cfg)
	assert.NoError(t, err)
	assert.NoError(t, os.WriteFile(tempFile, cfgData, 0644))

	params := DefaultParams(tempDir)
	cfgAfter, err := ParseBinary[Config](params)
	assert.NoError(t, err)
	assert.Equal(t, cfg, cfgAfter)

	params = DefaultParams()
	cfgAfter, err = ParseBinary[Config](params)
	assert.Nil(t, cfgAfter)
	assert.Error(t, err)
}
