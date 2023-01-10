package common_test

import (
	"dart/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {
	assert.True(t, common.FileExists("util_test.go"))
	assert.False(t, common.FileExists("sdlfkjleiilksdgls"))
}

func TestTestsAreRunning(t *testing.T) {
	assert.True(t, common.TestsAreRunning())
}
