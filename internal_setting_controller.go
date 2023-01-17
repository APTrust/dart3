package main

import (
	"dart/common"
)

func (a *App) InternalSettingList() *Response {
	response := NewResponse("Settings", "internal_setting/list.html")
	settings, err := common.InternalSettingList("obj_name", 1000, 0)
	if err != nil {
		response.Data["error"] = err.Error()
	}
	response.Data["items"] = settings
	return response.RenderContent()
}
