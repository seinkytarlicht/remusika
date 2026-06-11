package helper

import (
	"os"
	"path/filepath"
)

func CleanUpTemp() {
	files, _ := filepath.Glob("/tmp/remusika-tmpimg-*")
	for _, f := range files {
		os.Remove(f)
	}
}
