package application

import (
	"dart/common"
	"fmt"
)

func (a *App) AppSettingList() *Response {
	response := NewResponse("Settings", "app_setting/list.html")
	settings, err := common.AppSettingList("obj_name", 1000, 0)
	if err != nil {
		response.Data["error"] = err.Error()
	}
	response.Data["items"] = settings
	return response.RenderContent()
}

func (a *App) AppSettingCreate() *Response {
	setting := common.NewAppSetting("", "")
	response := NewResponse("Settings", "app_setting/form.html")
	form := setting.ToForm()
	response.Data["form"] = form
	return response.RenderContent()
}

func (a *App) AppSettingEdit(id string) *Response {
	response := NewResponse("Settings", "app_setting/form.html")
	setting, err := common.AppSettingFind(id)
	if err != nil {
		response.Data["error"] = err.Error()
	}
	form := setting.ToForm()
	form.UserCanDelete = setting.UserCanDelete
	response.Data["form"] = form
	return response.RenderContent()
}

func (a *App) AppSettingSave(setting *common.AppSetting) *Response {
	response := NewResponse("Settings", "app_setting/list.html")
	err := setting.Save()
	if err != nil {
		response.Data["error"] = err.Error()
		return response.RenderContent()
	} else {
		response.Data["flash"] = fmt.Sprintf("Saved setting %s", setting.Name)
	}
	settings, err := common.AppSettingList("obj_name", 1000, 0)
	if err != nil {
		response.Data["error"] = err.Error()
	}
	response.Data["items"] = settings
	return response.RenderContent()
}

func (a *App) AppSettingDelete(id string) *Response {
	response := NewResponse("Settings", "app_setting/list.html")
	setting, err := common.AppSettingFind(id)
	if err == nil {
		err = setting.Delete()
	}
	if err == nil {
		response.Data["flash"] = "Setting deleted."
	} else {
		response.Data["error"] = err.Error()
	}
	settings, err := common.AppSettingList("obj_name", 1000, 0)
	if err != nil {
		response.Data["error"] = err.Error()
	}
	response.Data["items"] = settings
	return response.RenderContent()
}
