package main

import (
	"dart/application"
	"dart/common"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

var app *application.App

func main() {
	// Create an instance of the app structure
	app = application.GetAppInstance()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "DART",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup:        app.Startup,
		OnShutdown:       app.Shutdown,
		Bind: []interface{}{
			app,
		},
		Logger:             logger.NewFileLogger(common.LogFilePath()),
		LogLevel:           logger.DEBUG,
		LogLevelProduction: logger.INFO,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
