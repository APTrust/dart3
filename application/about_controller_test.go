package application_test

import (
	"dart/application"
	"dart/common"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAboutShow(t *testing.T) {
	app := application.GetAppInstance()
	resp := app.AboutShow()
	require.NotNil(t, resp)
	require.NotEmpty(t, resp.Content)
	assert.Contains(t, resp.Content, common.DataFilePath())
	assert.Contains(t, resp.Content, common.LogFilePath())
}
