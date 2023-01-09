package main

import (
	"dart/common"
)

func (a *App) AppSettingList() Response {
	response := a.initResponse("Settings")
	response.Content = "App Setting List"
	response.Content = a.renderTemplate("app_setting/list.html", nil)
	return response
}

func (a *App) AppSettingCreate() Response {
	setting := common.NewAppSetting()
	response := a.initResponse("Settings")
	form := setting.ToForm()
	form.CancelFunction = "AppSettingList"
	form.SubmitFunction = "AppSettingSave"
	form.DeleteFunction = "AppSettingDelete"
	data := make(map[string]interface{})
	data["Form"] = form
	response.Content = a.renderTemplate("app_setting/form.html", data)
	return response
}

func (a *App) AppSettingEdit(id string) Response {
	response := a.initResponse("Settings")
	response.Content = "App Setting Edit"
	//response.Content = a.renderTemplate("dashboard/show.html", nil)
	return response
}

func (a *App) AppSettingSave() Response {
	response := a.initResponse("Settings")
	response.Content = "App Setting Save"
	//response.Content = a.renderTemplate("dashboard/show.html", nil)
	return response
}

func (a *App) AppSettingDelete(id string) Response {
	response := a.initResponse("Settings")
	response.Content = "App Setting Delete"
	//response.Content = a.renderTemplate("dashboard/show.html", nil)
	return response
}
