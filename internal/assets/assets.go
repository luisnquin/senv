package assets

import (
	_ "embed"
)

var (
	//go:embed templates/.env.tpl
	dotEnvTpl string
	//go:embed templates/help.tpl
	helpTpl string
)

func GetDotEnvTpl() string { return dotEnvTpl }

func GetHelpTpl() string { return helpTpl }
