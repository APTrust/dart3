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

func (a *App) DashboardShow() string {
	buf := bytes.Buffer{}
	err := a.Context.Templates.ExecuteTemplate(&buf, "dashboard/show.html", nil)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
