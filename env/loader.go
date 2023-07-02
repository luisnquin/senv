package env

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/luisnquin/senv/fsutils"
	"github.com/samber/lo"
)

// If returns a boolean wether is a YAML and/or .env files in the current working
// directory, it performs a folder tree traversal to above until find a .git folder
// if even in that case no single useful file is found then `false` value is returned.
func HasUsableWorkDir(fromDir string) bool {
	path := resolveUsableWorkDirectory(fromDir, false)

	return path != ""
}

func resolveUsableWorkDir(fromDir string) (string, error) {
	path := resolveUsableWorkDirectory(fromDir, false)
	if path == "" {
		return "", errors.New("work dir cannot be resolved")
	}

	return path, nil
}

func resolveUsableWorkDirectory(searchDir string, gitFolderFoundOnce bool) string {
	dirEntries, err := os.ReadDir(searchDir)
	if err != nil {
		return ""
	}

	for _, entry := range dirEntries {
		if entry.IsDir() {
			continue
		}

		fileName := entry.Name()

		if lo.Contains(configFiles, fileName) { // || strings.HasSuffix(fileName, ".env") {
			return searchDir
		}
	}

	if gitFolderFoundOnce {
		return ""
	}

	parentDir := filepath.Dir(searchDir)

	for {
		if fsutils.IsRootFolder(parentDir) {
			break
		}

		if fsutils.DirExists(filepath.Join(parentDir, ".git")) {
			return resolveUsableWorkDirectory(parentDir, true)
		}

		parentDir = filepath.Dir(parentDir)
	}

	return ""
}
