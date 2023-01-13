package main

import (
	"bytes"
	"dart/common"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Response struct {
	Content      string                 `json:"content"`
	ModalContent string                 `json:"modalContent"`
	Nav          string                 `json:"nav"`
	TemplateName string                 `json:"-"`
	Data         map[string]interface{} `json:"-"`
}

func NewResponse(section, templateName string) *Response {
	data := make(map[string]interface{})
	data["section"] = section // nav section
	return &Response{
		Data:         data,
		TemplateName: templateName,
	}
}

func (r *Response) RenderContent() *Response {
	runtime.LogDebugf(app.ctx, "Rendering content '%s' with data %v", r.TemplateName, r.Data)
	r.Content = r.renderTemplate(r.TemplateName)
	r.Nav = r.renderTemplate("partials/nav.html")
	return r
}

func (r *Response) RenderModal() *Response {
	runtime.LogDebugf(app.ctx, "Rendering modal '%s' with data %v", r.TemplateName, r.Data)
	r.ModalContent = r.renderTemplate(r.TemplateName)
	return r
}

func (r *Response) renderTemplate(name string) string {
	buf := bytes.Buffer{}
	err := common.Dart.Templates.ExecuteTemplate(&buf, name, r.Data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
