package main

func (a *App) AppSettingList() Response {
	response := a.initResponse("Settings")
	response.Content = "App Setting List"
	response.Content = a.renderTemplate("app_setting/list.html", nil)
	return response
}

func (a *App) AppSettingEdit() Response {
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
