package common

import (
	"html/template"
)

type Context struct {
	Templates *template.Template
	// DB
	// Log
}

func NewContext() *Context {
	return &Context{
		Templates: template.Must(template.New("").Funcs(getFuncMap()).ParseGlob("templates/**/*.html")),
	}
}

func getFuncMap() template.FuncMap {
	return template.FuncMap{
		"dateISO":       DateISO,
		"dateTimeISO":   DateTimeISO,
		"dateUS":        DateUS,
		"dateTimeUS":    DateTimeUS,
		"defaultString": DefaultString,
		"unixToISO":     UnixToISO,
		"yesNo":         YesNo,
	}
}
