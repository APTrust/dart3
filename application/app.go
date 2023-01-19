package application

import (
	"context"
	"dart/common"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	go_runtime "runtime"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var app *App

// App struct
type App struct {
	ctx  context.Context
	Dart *common.DartContext
}

// GetAppInstance returns the App struct, which is a singleton.
func GetAppInstance() *App {
	if app == nil {
		app = &App{}
	}
	return app
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.Dart = common.Dart
}

func (a *App) Shutdown(ctx context.Context) {
	runtime.LogDebug(app.ctx, "Shutting down.")
}

func (a *App) DashboardShow() *Response {
	response := NewResponse("Dashboard", "dashboard/show.html")
	return response.RenderContent()
}

// OpenExternal opens a link in an external browser.
func (a *App) OpenExternal(_url string) *Response {
	var cmd string
	var args []string
	switch go_runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}

	msg := fmt.Sprintf("%s %v", cmd, args)
	runtime.LogDebug(a.ctx, msg)
	args = append(args, _url)
	exec.Command(cmd, args...).Start()
	return &Response{}
}

func (a *App) AboutShow() *Response {
	appPath := ""
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		appPath = fmt.Sprintf("Can't get path: %s", err.Error())
	} else {
		appPath, _ = filepath.Abs(file)
	}
	response := NewResponse("Help", "about/index.html")
	response.Data["version"] = "3.x-alpha"
	response.Data["appPath"] = appPath
	response.Data["userDataPath"] = common.DataFilePath()
	response.Data["logFilePath"] = common.LogFilePath()
	return response.RenderContent()
}
