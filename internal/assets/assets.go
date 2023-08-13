package assets

import "embed"

var (
	//go:embed templates/.env.tpl
	dotEnvTpl string
	//go:embed templates/help.tpl
	helpTpl string
	//go:embed completions
	completionsFolder embed.FS
	//go:embed config/senv.example.yaml
	configExample []byte
)

func GetCompletionsFolder() embed.FS { return completionsFolder }

func GetDotEnvTpl() string { return dotEnvTpl }

func GetHelpTpl() string { return helpTpl }

func GetExampleConfig() []byte { return configExample }
