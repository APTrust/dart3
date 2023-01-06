package main

import (
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
