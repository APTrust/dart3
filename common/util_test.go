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

func TestLooksLikeHyperTextURL(t *testing.T) {
	assert.True(t, common.LooksLikeHypertextURL("http://example.com/api"))
	assert.True(t, common.LooksLikeHypertextURL("http://localhost/api"))
	assert.True(t, common.LooksLikeHypertextURL("https://repo.example.com/api/v2"))
	assert.False(t, common.LooksLikeHypertextURL("ftp://example.com/upload"))
	assert.False(t, common.LooksLikeHypertextURL("ταὐτὰ παρίσταταί"))
	assert.False(t, common.LooksLikeHypertextURL(""))
	assert.False(t, common.LooksLikeHypertextURL("6"))
}
