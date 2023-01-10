package common

import (
	"errors"
	"os"
	"strings"
)

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
