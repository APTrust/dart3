package application

import (
	"dart/common"
	"fmt"
)

func (a *App) RemoteRepositoryList() *Response {
	response := NewResponse("Settings", "remote_repository/list.html")
	repos, err := common.RemoteRepositoryList("obj_name", 1000, 0)
	if err != nil {
		response.Data["error"] = err.Error()
	}
	response.Data["items"] = repos
	return response.RenderContent()
}

func (a *App) RemoteRepositoryCreate() *Response {
	repo := common.NewRemoteRepository()
	response := NewResponse("Settings", "remote_repository/form.html")
	form := repo.ToForm()
	response.Data["form"] = form
	return response.RenderContent()
}

func (a *App) RemoteRepositoryEdit(id string) *Response {
	response := NewResponse("Settings", "remote_repository/form.html")
	repo, err := common.RemoteRepositoryFind(id)
	if err != nil {
		response.Data["error"] = err.Error()
	}
	form := repo.ToForm()
	response.Data["form"] = form
	return response.RenderContent()
}

func (a *App) RemoteRepositorySave(repo *common.RemoteRepository) *Response {
	response := NewResponse("Settings", "remote_repository/list.html")
	err := repo.Save()
	if err != nil {
		response.TemplateName = "remote_repository/form.html"
		response.Data["form"] = repo.ToForm()
		response.Data["error"] = err.Error()
		return response.RenderContent()
	} else {
		response.Data["flash"] = fmt.Sprintf("Saved remote repository %s", repo.Name)
	}
	repos, err := common.RemoteRepositoryList("obj_name", 1000, 0)
	if err != nil {
		response.Data["error"] = err.Error()
	}
	response.Data["items"] = repos
	return response.RenderContent()
}

func (a *App) RemoteRepositoryDelete(id string) *Response {
	response := NewResponse("Settings", "remote_repository/list.html")
	repo, err := common.RemoteRepositoryFind(id)
	if err == nil {
		err = repo.Delete()
	}
	if err == nil {
		response.Data["flash"] = "Deleted remote repository " + repo.Name
	} else {
		response.Data["error"] = err.Error()
	}
	repos, err := common.RemoteRepositoryList("obj_name", 1000, 0)
	if err != nil {
		response.Data["error"] = err.Error()
	}
	response.Data["items"] = repos
	return response.RenderContent()
}
