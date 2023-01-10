package common

import (
	"database/sql"
	"html/template"
	"log"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

type Context struct {
	Templates *template.Template
	DB        *sql.DB
	Log       *log.Logger
	Paths     *Paths
}

func NewContext() *Context {
	paths := NewPaths()
	return &Context{
		Templates: initTemplates(),
		DB:        initDB(paths),
		Log:       initLogger(paths),
		Paths:     paths,
	}
}

func initLogger(paths *Paths) *log.Logger {
	logFile := path.Join(paths.LogDir, "dart.log")
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	return log.New(f, "", log.LstdFlags)
}

func initDB(paths *Paths) *sql.DB {
	dbPath := path.Join(paths.DataDir, "dart.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func initTemplates() *template.Template {
	return template.Must(template.New("").Funcs(getFuncMap()).ParseGlob("templates/**/*.html"))
}

func getFuncMap() template.FuncMap {
	return template.FuncMap{
		"dateISO":       DateISO,
		"dateTimeISO":   DateTimeISO,
		"dateUS":        DateUS,
		"dateTimeUS":    DateTimeUS,
		"defaultString": DefaultString,
		"dict":          Dict,
		"unixToISO":     UnixToISO,
		"yesNo":         YesNo,
	}
}
