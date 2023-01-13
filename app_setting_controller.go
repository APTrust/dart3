package main

import (
	"dart/common"
	"fmt"
)

// TODO: Clean up & dedupe code.
//       Find and abstract the common patterns.
//       Also decide on caps vs. lowercase field names in templates.

func (a *App) AppSettingList() Response {
	response := a.initResponse("Settings")
	settings, err := common.AppSettingList("obj_name", 1000, 0)
	if err != nil {
		response.Error += err.Error()
	}
	data := make(map[string]interface{})
	data["items"] = settings
	data["error"] = response.Error
	data["flash"] = response.Flash
	response.Content = a.renderTemplate("app_setting/list.html", data)
	return response
}

func (a *App) AppSettingCreate() Response {
	setting := common.NewAppSetting("", "")
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
	setting, err := common.AppSettingFind(id)
	if err != nil {
		response.Error = err.Error()
	}
	form := setting.ToForm()
	form.CancelFunction = "AppSettingList"
	form.SubmitFunction = "AppSettingSave"
	form.DeleteFunction = "AppSettingDelete"
	form.UserCanDelete = setting.UserCanDelete
	data := make(map[string]interface{})
	data["Form"] = form
	data["error"] = response.Error
	response.Content = a.renderTemplate("app_setting/form.html", data)
	return response
}

func (a *App) AppSettingSave(setting *common.AppSetting) Response {
	response := a.initResponse("Settings")
	//setting := &common.AppSetting{}
	//err := json.Unmarshal([]byte(jsonStr), setting)
	//if err == nil {
	err := setting.Save()
	//}
	if err != nil {
		response.Error = err.Error()
	} else {
		response.Flash = fmt.Sprintf("Saved setting %s", setting.Name)
	}
	settings, err := common.AppSettingList("obj_name", 1000, 0)
	if err != nil {
		response.Error += err.Error()
	}
	data := make(map[string]interface{})
	data["items"] = settings
	data["error"] = response.Error
	data["flash"] = response.Flash
	response.Content = a.renderTemplate("app_setting/list.html", data)
	return response
}

func (a *App) AppSettingDelete(id string) Response {
	response := a.initResponse("Settings")
	setting, err := common.AppSettingFind(id)
	if err == nil {
		err = setting.Delete()
	}
	if err == nil {
		response.Flash = "Setting deleted."
	} else {
		response.Error = err.Error()
	}
	settings, err := common.AppSettingList("obj_name", 1000, 0)
	if err != nil {
		response.Error += err.Error()
	}
	data := make(map[string]interface{})
	data["items"] = settings
	data["error"] = response.Error
	data["flash"] = response.Flash
	response.Content = a.renderTemplate("app_setting/list.html", data)
	return response
}
