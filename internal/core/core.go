package core

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/luisnquin/senv/internal/assets"
	"github.com/luisnquin/senv/internal/fsutils"
	"github.com/samber/lo"
)

// Possible config files.
func getConfigFiles() []string {
	return []string{"senv.yaml", "senv.yml"}
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
func GenerateDotEnv(e Environment, binds map[string]string, useExportPrefix bool) ([]byte, error) {
	var b bytes.Buffer

	dotEnvTpl := template.Must(template.New(".env").Parse(assets.GetDotEnvTpl()))

	variables, err := getParsedVariables(e.Variables, binds)
	if err != nil {
		return nil, err
	}

	data := map[string]any{
		"sourceName": e.Name,
		"variables":  variables,
		"useExport":  useExportPrefix,
	}

	if err := dotEnvTpl.Execute(&b, data); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func getParsedVariables(originalVars map[string]string, binds map[string]string) (map[string]string, error) {
	vars := make(map[string]string, len(originalVars))

	for name, value := range originalVars {
		t, err := template.New(name).Parse(value)
		if err != nil {
			return nil, err
		}

		var b bytes.Buffer

		if err := t.Execute(&b, binds); err != nil {
			return nil, fmt.Errorf("unable to parse %s: %w", name, err)
		}

		vars[name] = b.String()
	}

	return vars, nil
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
