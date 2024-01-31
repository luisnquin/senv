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
