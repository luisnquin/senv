package core

import (
	"bytes"
	_ "embed"
	"errors"
	"os"
	"path/filepath"
	"text/template"

	"github.com/luisnquin/senv/internal/assets"
	"github.com/luisnquin/senv/internal/fsutils"
	"github.com/samber/lo"
)

// Possible config files.
var configFiles = []string{"senv.yaml", "senv.yml"}

// Returns the content of an .env file generated from the given environment variables configuration.
//
// Example:
//
//	DB_HOST="localhost"
//	DB_USER="root"
//	DB_PASSWORD="no_password"
//
// Additionally, the `export` prefix could be added for each environment variable.
//
// Example:
//
//	export DB_HOST="localhost"
//	export DB_USER="root"
//	export DB_PASSWORD="no_password"
func GenerateDotEnv(e Environment, useExportPrefix bool) ([]byte, error) {
	var b bytes.Buffer

	data := map[string]any{
		"sourceName": e.Name,
		"variables":  e.Variables,
		"useExport":  useExportPrefix,
	}

	t := template.Must(template.New(".env").Parse(assets.GetDotEnvTpl()))

	if err := t.Execute(&b, data); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// If returns a boolean wether is a YAML and/or .env files in the current working
// directory, it performs a folder tree traversal to above until find a .git folder
// if even in that case no single useful file is found then `false` value is returned.
func WorkDirHasProgramFiles(fromDir string) bool {
	path := workDirHasProgramFiles(fromDir, false)

	return path != ""
}

func resolveUsableWorkDir(fromDir string) (string, error) {
	path := workDirHasProgramFiles(fromDir, false)
	if path == "" {
		return "", errors.New("work directory cannot be resolved")
	}

	return path, nil
}

func workDirHasProgramFiles(searchDir string, gitFolderFoundOnce bool) string {
	dirEntries, err := os.ReadDir(searchDir)
	if err != nil {
		return ""
	}

	for _, entry := range dirEntries {
		if entry.IsDir() {
			continue
		}

		fileName := entry.Name()

		if lo.Contains(configFiles, fileName) {
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
			return workDirHasProgramFiles(parentDir, true)
		}

		parentDir = filepath.Dir(parentDir)
	}

	return ""
}
