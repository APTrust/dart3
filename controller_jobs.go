package main

func (a *App) JobShowFiles() *Response {
	response := NewResponse("Jobs", "job/files.html")
	return response.RenderContent()
}
