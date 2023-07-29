package yacl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/andrew528i/yacl/env"
	"github.com/andrew528i/yacl/file"
	"github.com/andrew528i/yacl/flags"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestSetEnvPrefix(t *testing.T) {
	type Config struct {
		Hostname string
	}

	assert.NoError(t, os.Setenv("HOSTNAME", "test123"))
	cfg, err := env.Parse[Config](env.DefaultParams())
	assert.NoError(t, err)
	assert.Equal(t, "test123", cfg.Hostname)
	os.Clearenv()

	SetEnvPrefix("APP")
	assert.NoError(t, os.Setenv("APP_HOSTNAME", "test321"))
	cfg, err = env.Parse[Config](env.DefaultParams())
	assert.NoError(t, err)
	assert.Equal(t, "test321", cfg.Hostname)
	os.Clearenv()

	SetEnvPrefix("")
}

func TestSetEnvDelimiter(t *testing.T) {
	type Config struct {
		HTTPPort uint
	}

	assert.NoError(t, os.Setenv("HTTP_PORT", "8081"))
	cfg, err := env.Parse[Config](env.DefaultParams())
	assert.NoError(t, err)
	assert.Equal(t, uint(8081), cfg.HTTPPort)
	os.Clearenv()

	SetEnvDelimiter("__")
	assert.NoError(t, os.Setenv("HTTP__PORT", "8082"))
	cfg, err = env.Parse[Config](env.DefaultParams())
	assert.NoError(t, err)
	assert.Equal(t, uint(8082), cfg.HTTPPort)
	os.Clearenv()

	SetEnvDelimiter("_")
}

func TestSetFilename(t *testing.T) {
	type Config struct {
		Hostname string
	}

	tempDir := t.TempDir()

	SetFilename("my-config")

	yamlBytes, err := yaml.Marshal(&Config{Hostname: "localhost-test"})
	assert.NoError(t, err)
	assert.NoError(t, os.WriteFile(filepath.Join(tempDir, "my-config.yaml"), yamlBytes, 0644))

	cfg, err := file.ParseYAML[Config](file.DefaultParams(tempDir))
	assert.NoError(t, err)
	assert.Equal(t, "localhost-test", cfg.Hostname)

	SetFilename("config")
}

func TestSetFlagDelimiter(t *testing.T) {
	type Config struct {
		DatabaseHostname string
	}

	SetFlagDelimiter("_")

	normalArgs := os.Args
	os.Args = []string{
		"cmd",
		"-database_hostname", "hello",
	}

	cfg, err := flags.Parse[Config](flags.DefaultParams())
	assert.NoError(t, err)
	assert.Equal(t, "hello", cfg.DatabaseHostname)

	os.Args = normalArgs
	SetFlagDelimiter("-")
}
