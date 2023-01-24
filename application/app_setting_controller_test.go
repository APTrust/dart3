package application_test

import (
	"dart/application"
	"dart/common"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAppSettingController(t *testing.T) {
	defer common.ClearDartTable()
	app := application.GetAppInstance()

	testAppSettingCreate(t, app)
	testAppSettingSave(t, app)
	testAppSettingSaveInvalid(t, app)
	testAppSettingEdit(t, app)
	testAppSettingDelete(t, app)
	testAppSettingList(t, app)
}

func testAppSettingCreate(t *testing.T, app *application.App) {
	resp := app.AppSettingCreate()
	assert.NotEmpty(t, resp.Content)
	assert.Contains(t, resp.Content, `<form method="post" action="#" id="AppSetting">`)
	assert.Contains(t, resp.Content, `<input type="hidden" name="ID"`)
}

func testAppSettingSave(t *testing.T, app *application.App) {
	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("Name_%d", i)
		value := fmt.Sprintf("Value_%d", i)
		setting := common.NewAppSetting(name, value)
		resp := app.AppSettingSave(setting)
		assert.Nil(t, resp.Data["error"])

		savedSetting, err := common.AppSettingFind(setting.ID)
		assert.Nil(t, err)
		assert.NotNil(t, savedSetting)
	}
}

func testAppSettingSaveInvalid(t *testing.T, app *application.App) {
	setting := common.NewAppSetting("", "")
	resp := app.AppSettingSave(setting)
	assert.Equal(t, "object contains validation errors", resp.Data["error"])
	assert.Contains(t, resp.Content, "Name cannot be empty")
	assert.Contains(t, resp.Content, "Value cannot be empty")
}

func testAppSettingEdit(t *testing.T, app *application.App) {
	setting := getFirstAppSetting(t, app)
	resp := app.AppSettingEdit(setting.ID)
	require.Empty(t, resp.Data["error"])
	require.NotEmpty(t, resp.Content)
	assert.Contains(t, resp.Content, `<form method="post" action="#" id="AppSetting">`)
	assert.Contains(t, resp.Content, fmt.Sprintf(`<input type="hidden" name="ID" value="%s"`, setting.ID))
	assert.Contains(t, resp.Content, setting.Name)
	assert.Contains(t, resp.Content, setting.Value)
}

func testAppSettingDelete(t *testing.T, app *application.App) {
	setting := getFirstAppSetting(t, app)
	resp := app.AppSettingDelete(setting.ID)
	require.Empty(t, resp.Data["error"])
	require.NotEmpty(t, resp.Content)

	deletedSetting, err := common.AppSettingFind(setting.ID)
	assert.NotNil(t, err)
	assert.Nil(t, deletedSetting)
}

func testAppSettingList(t *testing.T, app *application.App) {
	// We added five settings above, then deleted one,
	// so we should have four left.
	resp := app.AppSettingList()
	require.Empty(t, resp.Data["error"])
	require.NotEmpty(t, resp.Content)
	assert.Contains(t, resp.Content, "Name_1")
	assert.Contains(t, resp.Content, "Name_2")
	assert.Contains(t, resp.Content, "Name_3")
	assert.Contains(t, resp.Content, "Name_4")
}

func getFirstAppSetting(t *testing.T, app *application.App) *common.AppSetting {
	settings, err := common.AppSettingList("obj_name", 1, 0)
	require.Nil(t, err)
	require.NotEmpty(t, settings)
	setting := settings[0]
	require.NotNil(t, setting)
	return setting
}
