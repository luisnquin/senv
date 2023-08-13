package assets

import _ "embed"

var (
	//go:embed .env.tpl
	dotEnvTpl string
	// go:embed helpTpl.tpl
	helpTpl string
)

func GetDotEnvTpl() string { return dotEnvTpl }
func GetHelpTpl() string   { return helpTpl }
