package common

import (
	"fmt"

	"github.com/google/uuid"
)

type InternalSetting struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func NewInternalSetting(name, value string) *InternalSetting {
	return &InternalSetting{
		ID:    uuid.NewString(),
		Name:  name,
		Value: value,
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
	return ObjSave(setting)
}

func (setting *InternalSetting) Delete() error {
	return ObjDelete(setting.ID)
}

func (setting *InternalSetting) String() string {
	return fmt.Sprintf("InternalSetting: '%s' = '%s'", setting.Name, setting.Value)
}
