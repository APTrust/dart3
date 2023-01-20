package common

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

var reUrl = regexp.MustCompile(`(^http://localhost)|(https?://(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,4}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*))`)

// FileExists returns true if a file or directory exists at filePath.
func FileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}

// TestsAreRunning returns true when code is running under "go test"
func TestsAreRunning() bool {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-test.") {
			return true
		}
	}
	return false
}

// LooksLikeHypertextURL returns true if str looks like an
// HTTP or HTTPS URL.
func LooksLikeHypertextURL(str string) bool {
	return reUrl.MatchString(str)
}
