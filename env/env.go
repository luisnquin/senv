package env

import (
	"bytes"
	_ "embed"
	"path/filepath"
	"text/template"
)

// An user defined environment.
type Environment struct {
	Name   string         `yaml:"name"`
	Values map[string]any `yaml:"values"`
}

//go:embed .env.tpl
var dotEnvTpl string

// Loads all the environments as possible, it doesn't not return an error
// in case of no results.
func LoadEnvironments(workDirPath string) ([]Environment, error) {
	workDirPath = filepath.Clean(workDirPath)

	config, err := loadConfig(workDirPath)
	if err != nil {
		return nil, err
	}

	for _, e := range config.Environments {
		for k, v := range config.Defaults {
			_, ok := e.Values[k]
			if !ok {
				e.Values[k] = v
			}
		}
	}

	return config.Environments, nil
}

func GenerateDotEnv(e Environment) ([]byte, error) {
	var b bytes.Buffer

	data := map[string]any{
		"sourceName": e.Name,
		"values":     e.Values,
	}

	t := template.Must(template.New(".env").Parse(dotEnvTpl))

	if err := t.Execute(&b, data); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
