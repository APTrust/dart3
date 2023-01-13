package main

import (
	"bytes"
	"dart/common"
)

type Response struct {
	Content      string                 `json:"content"`
	ModalContent string                 `json:"modalContent"`
	Nav          string                 `json:"nav"`
	Error        string                 `json:"error"`
	Flash        string                 `json:"flash"`
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
	r.Content = r.renderTemplate(r.TemplateName)
	r.Nav = r.renderTemplate("partials/nav.html")
	return r
}

func (r *Response) RenderModal() *Response {
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
