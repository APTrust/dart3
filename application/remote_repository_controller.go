package application

import (
	"dart/common"
)

func (a *App) RemoteRepositoryList() *Response {
	response := NewResponse("Settings", "remote_repository/list.html")

	return response.RenderContent()
}

func (a *App) RemoteRepositoryCreate() *Response {
	//repo := common.NewRemoteRepository()
	response := NewResponse("Settings", "remote_repository/form.html")
	response.Content = "This will be the remote repo form."

	return response.RenderContent()
}

func (a *App) RemoteRepositoryEdit(id string) *Response {
	response := NewResponse("Settings", "remote_repository/form.html")

	return response.RenderContent()
}

func (a *App) RemoteRepositorySave(repo *common.RemoteRepository) *Response {
	response := NewResponse("Settings", "remote_repository/list.html")

	return response.RenderContent()
}

func (a *App) RemoteRepositoryDelete(id string) *Response {
	response := NewResponse("Settings", "remote_repository/list.html")

	return response.RenderContent()
}
