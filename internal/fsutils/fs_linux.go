package fsutils

// Multi-platform way to check if the passed argument corresponds for the
// root directory or not.
func IsRootFolder(path string) bool {
	return path == "/"
}
