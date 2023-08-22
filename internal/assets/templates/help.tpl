{{.CommandName | green | bold }}{{if .Description}} - {{.Description | yellow }}{{end}}{{if .PrependMessage}}
{{.PrependMessage}}{{end}}
{{if .Subcommands}}
{{ "Available commands:" | magenta | underline }} {{range .Subcommands}}
  {{.LongName}}{{if .ShortName}} ({{.ShortName}}){{end}}{{if .Position}}{{if gt .Position 1}}  (position {{.Position}}){{end}}{{end}}{{if .Description}}   {{.Spacer}}{{.Description}}{{end}}{{end}}
{{end}}{{if (gt (len .Flags) 0)}}
{{ "Flags:" | magenta | underline }} {{if .Flags}}{{range .Flags}}
  {{if .ShortName}}-{{.ShortName | italic}} {{else}}   {{end}}{{if .LongName}}--{{.LongName | italic}}{{end}}{{if .Description}}   {{.Spacer}}{{.Description}}{{if .DefaultValue}} (default: {{.DefaultValue}}){{end}}{{end}}{{end}}{{end}}
{{end}}{{if .AppendMessage}}{{.AppendMessage}}
{{end}}{{if .Message}}
{{.Message}}{{end}}