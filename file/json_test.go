package file

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseJSON(t *testing.T) {
	type DatabaseConfig struct {
		Port    uint64 `json:"port"`
		Restart bool   `json:"restart"`
	}

	type LoggerConfig struct {
		Levels []uint `json:"levels"`
	}

	type Config struct {
		Username        string         `json:"username"`
		MaxTemperatures []float64      `json:"max_temperatures"`
		Database        DatabaseConfig `json:"database"`
		Logger          LoggerConfig   `json:"logger"`
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

	tempFile := filepath.Join(tempDir, "config.json")
	cfgData, err := json.Marshal(cfg)
	assert.NoError(t, err)
	assert.NoError(t, os.WriteFile(tempFile, cfgData, 0644))

	params := DefaultParams(tempDir)
	cfgAfter, err := ParseJSON[Config](params)
	assert.NoError(t, err)
	assert.Equal(t, cfg, cfgAfter)

	params = DefaultParams()
	cfgAfter, err = ParseJSON[Config](params)
	assert.Nil(t, cfgAfter)
	assert.Error(t, err)
}
