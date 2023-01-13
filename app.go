package main

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

// App struct
type App struct {
	ctx  context.Context
	Dart *common.DartContext
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.Dart = common.Dart
}

func (a *App) DashboardShow() *Response {
	response := NewResponse("Dashboard", "dashboard/show.html")
	return response.RenderContent()
}

// OpenExternal opens a link in an external browser.
func (a *App) OpenExternal(_url string) *Response {
	msg := fmt.Sprintf("Opening external link %s", _url)
	runtime.LogDebug(a.ctx, msg)

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

	msg = fmt.Sprintf("%s %v", cmd, args)
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
	// runtime.LogDebug(a.ctx, "~~~~~~~~~~~~ AboutShow ~~~~~~~~~~~~~~~~")
	response := NewResponse("Help", "about/index.html")
	response.Data["version"] = "3.x-alpha"
	response.Data["appPath"] = appPath
	response.Data["userDataPath"] = common.DataFilePath()
	response.Data["logFilePath"] = common.LogFilePath()
	return response.RenderContent()
}

// func (a *App) initResponse(section string) Response {
// 	return Response{
// 		Nav: a.renderNav(section),
// 	}
// }

// func (a *App) renderTemplate(name string, data interface{}) string {
// 	buf := bytes.Buffer{}
// 	err := a.Dart.Templates.ExecuteTemplate(&buf, name, data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return buf.String()
// }

// func (a *App) renderNav(section string) string {
// 	data := map[string]string{
// 		"section": section,
// 	}
// 	return a.renderTemplate("partials/nav.html", data)
// }
