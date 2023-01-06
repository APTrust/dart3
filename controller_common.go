package main

import "bytes"

func (a *App) DashboardShow() Response {
	response := a.initResponse("Dashboard")
	response.Content = a.renderTemplate("dashboard/show.html", nil)
	return response
}

func (a *App) initResponse(section string) Response {
	return Response{
		Nav: a.renderNav(section),
	}
}

func (a *App) renderTemplate(name string, data interface{}) string {
	buf := bytes.Buffer{}
	err := a.Context.Templates.ExecuteTemplate(&buf, name, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (a *App) renderNav(section string) string {
	data := map[string]string{
		"section": section,
	}
	return a.renderTemplate("partials/nav.html", data)
}
