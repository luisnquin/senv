package fsutils

import "os"

func FileExists(path string) bool {
	info, err := os.Stat(path)

	return err == nil && !info.IsDir()
}

func DirExists(path string) bool {
	info, err := os.Stat(path)

	return err == nil && info.IsDir()
}

// Multi-platform way to check if the passed argument corresponds for the
// root directory or not.
func IsRootFolder(path string) bool {
	return path == "/"
}
