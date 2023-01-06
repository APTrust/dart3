package main

import (
	"bytes"
	"context"
	"dart/common"
)

// App struct
type App struct {
	ctx     context.Context
	Context *common.Context
}

type Response struct {
	Content      string `json:"content"`
	ModalContent string `json:"modalContent"`
	Nav          string `json:"nav"`
	Error        string `json:"error"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.Context = common.NewContext()
}

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
