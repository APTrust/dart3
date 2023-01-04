package main

import (
	"bytes"
	"context"
	"dart/common"
	"fmt"
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

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) DashboardShow() Response {
	buf := bytes.Buffer{}
	err := a.Context.Templates.ExecuteTemplate(&buf, "dashboard/show.html", nil)
	if err != nil {
		panic(err)
	}
	return Response{
		Content: buf.String(),
		Nav:     a.RenderNav("Dashboard"),
	}
}

func (a *App) RenderNav(section string) string {
	data := map[string]string{
		"section": section,
	}
	buf := bytes.Buffer{}
	err := a.Context.Templates.ExecuteTemplate(&buf, "partials/nav.html", data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
