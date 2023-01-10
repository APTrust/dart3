package common

type InternalSetting struct {
	ID    string
	Name  string
	Value string
}

func (setting *InternalSetting) ObjID() string {
	return setting.ID
}

func (setting *InternalSetting) ObjName() string {
	return setting.Name
}

func (setting *InternalSetting) ObjType() string {
	return "InternalSetting"
}
