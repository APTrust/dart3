package common

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// AppSetting represents an application-wide setting that can be
// configured by the user. For example, the bagging directory
// into which DART writes new bags.
//
// Field names for JSON serialization match the old DART 2 names,
// so we don't break legacy installations.
type AppSetting struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	Value         string            `json:"value"`
	Help          string            `json:"help"`
	Errors        map[string]string `json:"errors"`
	UserCanDelete bool              `json:"userCanDelete"`
}

// NewAppSetting creates a new AppSetting with the specified name
// and value. UserCanDelete will be true by default. If a setting
// is required for DART to function properly (such as the Bagging
// Directory setting), set UserCanDelete to false.
func NewAppSetting(name, value string) *AppSetting {
	return &AppSetting{
		ID:            uuid.NewString(),
		Name:          name,
		Value:         value,
		UserCanDelete: true,
		Errors:        make(map[string]string),
	}
}

// AppSettingFind returns the AppSetting with the specified UUID,
// or sql.ErrNoRows if no matching record exists.
func AppSettingFind(uuid string) (*AppSetting, error) {
	result, err := ObjFind(uuid)
	if err != nil {
		return nil, err
	}
	return result.AppSetting, err
}

// AppSettingList returns a list of AppSettings with the specified
// order, offset and limit.
func AppSettingList(orderBy string, limit, offset int) ([]*AppSetting, error) {
	result, err := ObjList(TypeAppSetting, orderBy, limit, offset)
	if err != nil {
		return nil, err
	}
	return result.AppSettings, err
}

// ObjID returns this setting's object id (uuid).
func (setting *AppSetting) ObjID() string {
	return setting.ID
}

// ObjName returns this object's name, so names will be
// searchable and sortable in the DB.
func (setting *AppSetting) ObjName() string {
	return setting.Name
}

// ObjType returns this object's type name.
func (setting *AppSetting) ObjType() string {
	return TypeAppSetting
}

// Save saves this setting, if it determines the setting is valid.
// It returns common.ErrObjecValidation if the setting is invalid.
// Check setting.Errors if you get a validation error.
func (setting *AppSetting) Save() error {
	if !setting.Validate() {
		return ErrObjecValidation
	}
	return ObjSave(setting)
}

// Delete deletes this AppSetting. If the setting is marked with
// UserCanDelete = false, you'll get a common.ErrNotDeletable error.
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

// ToForm returns a form so the user can edit this AppSetting.
// The form can be rendered by the app_setting/form.html template.
func (setting *AppSetting) ToForm() *Form {
	form := NewForm(TypeAppSetting, setting.ID)

	form.AddField("ID", "ID", setting.ID, true)

	userCanDeleteField := form.AddField("UserCanDelete", "UserCanDelete", strconv.FormatBool(setting.UserCanDelete), true)
	userCanDeleteField.Cast = CastToBool

	nameField := form.AddField("Name", "Name", setting.Name, true)
	nameField.Error = setting.Errors["Name"]
	// If user cannot delete this field, they can't rename it either.
	// Renaming the setting would prevent the app from finding it,
	// an in the case of a required setting like "Bagging Directory,"
	// that would cause lots of problems.
	if !setting.UserCanDelete {
		nameField.Attrs["disabled"] = "disabled"
	}

	valueField := form.AddField("Value", "Value", setting.Value, true)
	valueField.Error = setting.Errors["Value"]
	valueField.Help = "If the setting has help text, it will be displayed here." // setting.Help

	return form
}

// Validate validates this setting, returning true if it's valid,
// false if not. If false, this sets specific error messages in the
// Errors map, which are suitable for display on the form.
func (setting *AppSetting) Validate() bool {
	setting.Errors = make(map[string]string)
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
