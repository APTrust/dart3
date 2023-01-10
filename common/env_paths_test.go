package common_test

import (
	"dart/common"
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPaths(t *testing.T) {
	p := common.NewPaths()
	require.NotNil(t, p)

	homeDir, _ := os.UserHomeDir()
	assert.Contains(t, p.DataDir, homeDir)
	assert.Contains(t, p.ConfigDir, homeDir)
	assert.Contains(t, p.CacheDir, homeDir)
	assert.Contains(t, p.HomeDir, homeDir)
	assert.Contains(t, p.LogDir, homeDir)
	assert.Contains(t, p.TempDir, common.AppName)

	switch runtime.GOOS {
	case "darwin":
		testMacOsPaths(t, p)
	case "windows":
		testWindowsPaths(t, p)
	case "linux":
		testLinuxPaths(t, p)
	}
}

func testMacOsPaths(t *testing.T, p *common.Paths) {
	assert.Contains(t, p.DataDir, "Library/Application Support")
	assert.Contains(t, p.ConfigDir, "Library/Preferences")
	assert.Contains(t, p.CacheDir, "Library/Caches")
	assert.Contains(t, p.LogDir, "Library/Logs")
}

func testWindowsPaths(t *testing.T, p *common.Paths) {
	assert.Contains(t, p.DataDir, "Data")
	assert.Contains(t, p.ConfigDir, "Config")
	assert.Contains(t, p.CacheDir, "Cache")
	assert.Contains(t, p.LogDir, "Log")
}

func testLinuxPaths(t *testing.T, p *common.Paths) {
	assert.Contains(t, p.DataDir, common.AppName)
	assert.Contains(t, p.ConfigDir, common.AppName)
	assert.Contains(t, p.CacheDir, common.AppName)
	assert.Contains(t, p.LogDir, common.AppName)
}
