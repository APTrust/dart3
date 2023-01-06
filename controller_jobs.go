package main

func (a *App) JobShowFiles() Response {
	response := a.initResponse("Jobs")
	response.Content = a.renderTemplate("job/files.html", nil)
	return response
}
