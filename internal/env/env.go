package env

import (
	"bytes"
	_ "embed"
	"text/template"
)

// An user defined environment.
type Environment struct {
	Name      string         `yaml:"name"`
	Variables map[string]any `yaml:"variables"`
}

//go:embed .env.tpl
var dotEnvTpl string

func GenerateDotEnv(e Environment, useExportPrefix bool) ([]byte, error) {
	var b bytes.Buffer

	data := map[string]any{
		"sourceName": e.Name,
		"variables":  e.Variables,
		"useExport":  useExportPrefix,
	}

	t := template.Must(template.New(".env").Parse(dotEnvTpl))

	if err := t.Execute(&b, data); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
