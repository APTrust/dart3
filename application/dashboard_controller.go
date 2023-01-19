package application

func (a *App) DashboardShow() *Response {
	response := NewResponse("Dashboard", "dashboard/show.html")
	return response.RenderContent()
}
