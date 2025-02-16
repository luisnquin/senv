package core

import (
	"bytes"
	_ "embed"
	"os"
	"path/filepath"
	"text/template"

	"github.com/luisnquin/senv/internal/assets"
	"github.com/luisnquin/senv/internal/fsutils"
	"github.com/samber/lo"
)

// Possible config files.
func getConfigFiles() []string {
	return []string{".senv", "senv.yaml", "senv.yml"}
}

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
func GenerateDotEnv(name string, groupedVariables []map[string]any, useExportPrefix bool) ([]byte, error) {
	var b bytes.Buffer

	data := map[string]any{
		"source_name":       name,
		"grouped_variables": groupedVariables,
		"use_export":        useExportPrefix,
	}

	t := template.Must(template.New(".env").Parse(assets.GetDotEnvTpl()))

	if err := t.Execute(&b, data); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// Returns a boolean if the current of a parent directory contains config files
// like: senv.yaml or senv.yml.
func HasConfigFiles(fromDir string) bool {
	path := getDirWithConfig(fromDir)

	return path != ""
}

// Tries to find senv.yaml or senv.yml files in the current directory or parents.
//
// If is the root directory and no config file is found then it returns an error.
func ResolveDirWithConfig(fromDir string) (string, error) {
	path := getDirWithConfig(fromDir)
	if path == "" {
		return "", getErrUnableToResolveWorkDir()
	}

	return path, nil
}

// Tries to find senv.yaml or senv.yml files in the current directory or parents.
//
// If is the root directory and no config file is found then it an empty string
// is returned.
func getDirWithConfig(searchDir string) string {
	dirEntries, err := os.ReadDir(searchDir)
	if err != nil {
		return ""
	}

	for _, entry := range dirEntries {
		if entry.IsDir() {
			continue
		}

		fileName := entry.Name()

		if lo.Contains(getConfigFiles(), fileName) {
			return searchDir
		}
	}

	parentDir := filepath.Dir(searchDir)

	for {
		if fsutils.IsRootFolder(parentDir) {
			break
		}

		for _, cf := range getConfigFiles() {
			if fsutils.FileExists(filepath.Join(parentDir, cf)) {
				return parentDir
			}
		}

		parentDir = filepath.Dir(parentDir)
	}

	return ""
}
