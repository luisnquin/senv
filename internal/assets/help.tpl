{{.CommandName}}{{if .Description}} - {{.Description}}{{end}}{{if .PrependMessage}}
{{.PrependMessage}}{{end}}
{{if .Subcommands}}
Available commands: {{range .Subcommands}}
  {{.LongName}}{{if .ShortName}} ({{.ShortName}}){{end}}{{if .Position}}{{if gt .Position 1}}  (position {{.Position}}){{end}}{{end}}{{if .Description}}   {{.Spacer}}{{.Description}}{{end}}{{end}}
{{end}}{{if (gt (len .Flags) 0)}}
Flags: {{if .Flags}}{{range .Flags}}
  {{if .ShortName}}-{{.ShortName}} {{else}}   {{end}}{{if .LongName}}--{{.LongName}}{{end}}{{if .Description}}   {{.Spacer}}{{.Description}}{{if .DefaultValue}} (default: {{.DefaultValue}}){{end}}{{end}}{{end}}{{end}}
{{end}}{{if .AppendMessage}}{{.AppendMessage}}
{{end}}{{if .Message}}
{{.Message}}{{end}}