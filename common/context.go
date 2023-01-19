package common

import (
	"database/sql"
	"html/template"
	"log"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

var Dart *DartContext

type DartContext struct {
	Templates *template.Template
	DB        *sql.DB
	Log       *Logger
	Paths     *Paths
}

func init() {
	paths := NewPaths()
	Dart = &DartContext{
		Templates: initTemplates(),
		DB:        initDB(paths),
		Paths:     paths,
		Log:       initLogger(LevelDebug),
	}
	InitSchema()
}

func initDB(paths *Paths) *sql.DB {
	db, err := sql.Open("sqlite3", DataFilePath())
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func initTemplates() *template.Template {
	// Note: When running `wails dev`, we'll load templates from the "templates" dir.
	// When running `go test ./...`, go may descend into subdirectories, so we have
	// to look up the directory tree to find the templates; otherwise, we get a panic.
	var t *template.Template
	if FileExists("templates") {
		t = template.Must(template.New("").Funcs(getFuncMap()).ParseGlob("templates/**/*.html"))
	} else if FileExists("../templates") {
		t = template.Must(template.New("").Funcs(getFuncMap()).ParseGlob("../templates/**/*.html"))
	} else if FileExists("../../templates") {
		t = template.Must(template.New("").Funcs(getFuncMap()).ParseGlob("../../templates/**/*.html"))
	}
	return t
}

func LogFilePath() string {
	paths := NewPaths()
	return path.Join(paths.LogDir, "dart.log")
}

func DataFilePath() string {
	paths := NewPaths()
	dbPath := path.Join(paths.DataDir, "dart.db")
	// Run tests in an in-memory db, so we don't pollute
	// our actual dart db.
	if TestsAreRunning() {
		dbPath = ":memory:"
		//dbPath = path.Join(paths.HomeDir, "Desktop", "dart.db")
	}
	return dbPath
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
