package common

import (
	"database/sql"
	"html/template"
	"log"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

var Dart *DartContext

type DartContext struct {
	Templates *template.Template
	DB        *sql.DB
	Log       *log.Logger
	Paths     *Paths
}

func init() {
	paths := NewPaths()
	Dart = &DartContext{
		Templates: initTemplates(),
		DB:        initDB(paths),
		Log:       initLogger(paths),
		Paths:     paths,
	}
	InitSchema()
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
	// TODO: Relative paths are different between `wails dev` and `go test ./...`.
	// The latter won't find the templates. We need to fix that.
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
