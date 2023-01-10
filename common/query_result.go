package common

type QueryResult struct {
	AppSetting       *AppSetting
	AppSettings      []*AppSetting
	InternalSetting  *InternalSetting
	InternalSettings []*InternalSetting
	ObjType          string
}

func NewQueryResult(objType string) *QueryResult {
	return &QueryResult{
		ObjType: objType,
	}
}

// func NewQueryResultSingle(obj interface{}, objType string) *QueryResult {
// 	list := make([]interface{}, 1)
// 	list[0] = obj
// 	return &QueryResult{
// 		items:   list,
// 		ObjType: objType,
// 		Count:   len(list),
// 	}
// }

// func (qr *QueryResult) AppSetting() *AppSetting {
// 	return qr.items[0].(*AppSetting)
// }

// func (qr *QueryResult) AppSettingList() []*AppSetting {
// 	list := make([]*AppSetting, len(qr.items))
// 	for i := 0; i < qr.items; i++ {
// 		list[i] = qr.items[i].(*AppSetting)
// 	}
// 	return list
// }

// func (qr *QueryResult) InternalSetting() *InternalSetting {
// 	return qr.items[0].(*InternalSetting)
// }

// func (qr *QueryResult) InternalSettingList() []*InternalSetting {
// 	list := make([]*InternalSetting, len(qr.items))
// 	for i := 0; i < qr.items; i++ {
// 		list[i] = qr.items[i].(*InternalSetting)
// 	}
// 	return list
// }
