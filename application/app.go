package application

import (
	"context"
	"dart/common"
	"fmt"
	"os/exec"
	go_runtime "runtime"
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
		app = &App{
			Dart: common.Dart,
		}
	}
	return app
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Shutdown(ctx context.Context) {
	a.Dart.Log.Debug("Shutting down.")
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
	a.Dart.Log.Debug(msg)
	args = append(args, _url)
	exec.Command(cmd, args...).Start()
	return &Response{}
}
