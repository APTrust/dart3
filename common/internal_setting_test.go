package common_test

import (
	"dart/common"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInternalSettingPersistence(t *testing.T) {
	//defer common.ClearDartTable()
	s1 := common.NewInternalSetting("Setting 1", "Value 1")
	s2 := common.NewInternalSetting("Setting 2", "Value 2")
	s3 := common.NewInternalSetting("Setting 3", "Value 3")
	assert.Nil(t, s1.Save())
	assert.Nil(t, s2.Save())
	assert.Nil(t, s3.Save())

	s1Reload, err := common.InternalSettingFind(s1.ID)
	require.Nil(t, err)
	require.NotNil(t, s1Reload)
	assert.Equal(t, s1.ID, s1Reload.ID)
	assert.Equal(t, s1.Name, s1Reload.Name)
	assert.Equal(t, s1.Value, s1Reload.Value)

	// TODO: List tests are failing :(
	settings, err := common.InternalSettingList("obj_name", 1, 0)
	require.Nil(t, err)
	require.Equal(t, 1, len(settings))
	assert.Equal(t, s1.ID, settings[0].ID)

	settings, err = common.InternalSettingList("obj_name", 100, 0)
	require.Nil(t, err)
	require.Equal(t, 3, len(settings))
	assert.Equal(t, s1.ID, settings[0].ID)
	assert.Equal(t, s2.ID, settings[1].ID)
	assert.Equal(t, s3.ID, settings[2].ID)
}
