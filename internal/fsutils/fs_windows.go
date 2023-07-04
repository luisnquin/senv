package fsutils

import (
	"os"
	"path/filepath"
)

// Multi-platform way to check if the passed argument corresponds for the
// root directory or not.
func IsRootFolder(path string) bool {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	rootFolder := filepath.VolumeName(workingDir) + "\\"

	return rootFolder == filepath.Clean(path)
}
