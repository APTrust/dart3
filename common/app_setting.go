package common

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type AppSetting struct {
	ID            string
	Name          string
	Value         string
	Help          string
	Errors        map[string]string
	UserCanDelete bool
}

func NewAppSetting() *AppSetting {
	return &AppSetting{
		ID:     uuid.NewString(),
		Errors: make(map[string]string),
	}
}

func AppSettingFromJson(jsonStr string) (*AppSetting, error) {
	setting := &AppSetting{}
	err := json.Unmarshal([]byte(jsonStr), setting)
	if err == nil {
		setting.Name = strings.TrimSpace(setting.Name)
		setting.Value = strings.TrimSpace(setting.Value)
	}
	return setting, err
}

func (setting *AppSetting) ToForm() *Form {
	form := NewForm("AppSetting")

	form.AddField("ID", "ID", setting.ID, true)
	form.AddField("UserCanDelete", "UserCanDelete", strconv.FormatBool(setting.UserCanDelete), true)
	nameField := form.AddField("Name", "Name", setting.Name, true)
	nameField.Error = setting.Errors["Name"]

	valueField := form.AddField("Value", "Value", setting.Value, true)
	valueField.Error = setting.Errors["Value"]

	return form
}

func (setting *AppSetting) Validate() bool {
	isValid := true
	if setting.Name == "" {
		setting.Errors["Name"] = "Name cannot be empty."
		isValid = false
	}
	if setting.Value == "" {
		setting.Errors["Value"] = "Value cannot be empty."
		isValid = false
	}
	return isValid
}
