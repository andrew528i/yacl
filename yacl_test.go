package yacl

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYACL_SetEnvPrefix(t *testing.T) {
	type Config struct {
		Hostname string
	}

	assert.NoError(t, os.Setenv("APP_HOSTNAME", "localhost-321"))
	assert.NoError(t, os.Setenv("HOSTNAME", "localhost-123"))

	first := New[Config]()
	first.SetIgnoreFlags(true)
	first.SetEnvPrefix("APP")
	firstCfg, err := first.Parse()
	assert.NoError(t, err)
	assert.Equal(t, "localhost-321", firstCfg.Hostname)

	second := New[Config]()
	second.SetIgnoreFlags(true)
	secondCfg, err := second.Parse()
	assert.NoError(t, err)
	assert.Equal(t, "localhost-123", secondCfg.Hostname)
}
