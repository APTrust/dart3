package common

import "encoding/json"

type QueryResult struct {
	dbObjects []*DBObject
}

func NewQueryResultMultiple(list []*DBObject) *QueryResult {
	return &QueryResult{
		dbObjects: list,
	}
}

func NewQueryResultSingle(obj *DBObject) *QueryResult {
	list := make([]*DBObject, 1)
	list[0] = obj
	return &QueryResult{
		dbObjects: list,
	}
}

func (qr *QueryResult) AppSetting() (*AppSetting, error) {
	dbObj := qr.dbObjects[0]
	setting := &AppSetting{}
	err := json.Unmarshal([]byte(dbObj.Json), setting)
	return setting, err
}

func (qr *QueryResult) AppSettingList() ([]*AppSetting, error) {
	list := make([]*AppSetting, len(qr.dbObjects))
	for i := 0; i < len(qr.dbObjects); i++ {
		setting := &AppSetting{}
		err := json.Unmarshal([]byte(qr.dbObjects[i].Json), setting)
		if err != nil {
			Dart.Log.Printf("[ERROR] Json unmarshal error on AppSetting %s: %v", qr.dbObjects[i].Name, err)
			return nil, err
		}
		list[i] = setting
	}
	return list, nil
}
