package application_test

import (
	"dart/application"
	"dart/common"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInternalSettingList(t *testing.T) {
	defer common.ClearDartTable()
	require.Nil(t, common.NewInternalSetting("List Setting One", "Val One").Save())
	require.Nil(t, common.NewInternalSetting("List Setting Two", "Val Two").Save())

	app := application.GetAppInstance()
	resp := app.InternalSettingList()
	require.NotNil(t, resp)
	require.NotEmpty(t, resp.Content)
	assert.Contains(t, resp.Content, "List Setting One")
	assert.Contains(t, resp.Content, "Val One")
	assert.Contains(t, resp.Content, "List Setting Two")
	assert.Contains(t, resp.Content, "Val Two")
}
