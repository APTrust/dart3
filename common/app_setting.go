package common

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type AppSetting struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	Value         string            `json:"value"`
	Help          string            `json:"help"`
	Errors        map[string]string `json:"errors"`
	UserCanDelete bool              `json:"userCanDelete"`
}

func NewAppSetting() *AppSetting {
	return &AppSetting{
		ID:     uuid.NewString(),
		Errors: make(map[string]string),
	}
}

func AppSettingFind(uuid string) (*AppSetting, error) {
	result, err := ObjFind(uuid)
	if err != nil {
		return nil, err
	}
	return result.AppSetting, err
}

func AppSettingList(orderBy string, limit, offset int) ([]*AppSetting, error) {
	result, err := ObjList(TypeAppSetting, orderBy, limit, offset)
	if err != nil {
		return nil, err
	}
	return result.AppSettings, err
}

func (setting *AppSetting) ObjID() string {
	return setting.ID
}

func (setting *AppSetting) ObjName() string {
	return setting.Name
}

func (setting *AppSetting) ObjType() string {
	return TypeAppSetting
}

func (setting *AppSetting) Save() error {
	if !setting.Validate() {
		return ErrObjecValidation
	}
	return ObjSave(setting)
}

func (setting *AppSetting) Delete() error {
	if !setting.UserCanDelete {
		return ErrNotDeletable
	}
	return ObjDelete(setting.ID)
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
	form := NewForm(TypeAppSetting)

	form.AddField("ID", "ID", setting.ID, true)
	form.AddField("UserCanDelete", "UserCanDelete", strconv.FormatBool(setting.UserCanDelete), true)
	nameField := form.AddField("Name", "Name", setting.Name, true)
	nameField.Error = setting.Errors["Name"]

	valueField := form.AddField("Value", "Value", setting.Value, true)
	valueField.Error = setting.Errors["Value"]
	valueField.Help = "This help text should appear in a popup." // setting.Help

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
