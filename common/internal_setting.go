package common

import (
	"fmt"

	"github.com/google/uuid"
)

// InternalSetting is set by DART and cannot be edited by user.
// These settings may record when migrations were run, or other
// internal info. These settings cannot be created or edited by
// users.
type InternalSetting struct {
	ID     string            `json:"id"`
	Name   string            `json:"name"`
	Value  string            `json:"value"`
	Errors map[string]string `json:"errors"`
}

func NewInternalSetting(name, value string) *InternalSetting {
	return &InternalSetting{
		ID:     uuid.NewString(),
		Name:   name,
		Value:  value,
		Errors: make(map[string]string),
	}
}

func InternalSettingFind(uuid string) (*InternalSetting, error) {
	result, err := ObjFind(uuid)
	if err != nil {
		return nil, err
	}
	return result.InternalSetting, err
}

func InternalSettingList(orderBy string, limit, offset int) ([]*InternalSetting, error) {
	result, err := ObjList(TypeInternalSetting, orderBy, limit, offset)
	if err != nil {
		return nil, err
	}
	return result.InternalSettings, err
}

func (setting *InternalSetting) ObjID() string {
	return setting.ID
}

func (setting *InternalSetting) ObjName() string {
	return setting.Name
}

func (setting *InternalSetting) ObjType() string {
	return TypeInternalSetting
}

func (setting *InternalSetting) Save() error {
	if !setting.Validate() {
		return ErrObjecValidation
	}
	return ObjSave(setting)
}

func (setting *InternalSetting) Delete() error {
	return ObjDelete(setting.ID)
}

func (setting *InternalSetting) String() string {
	return fmt.Sprintf("InternalSetting: '%s' = '%s'", setting.Name, setting.Value)
}

// Validate validates this setting, returning true if it's valid,
// false if not. If false, this sets specific error messages in the
// Errors map, which are suitable for display on the form.
func (setting *InternalSetting) Validate() bool {
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
