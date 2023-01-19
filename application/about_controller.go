package application

import (
	"dart/common"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

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
