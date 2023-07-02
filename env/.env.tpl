# source: {{ .sourceName }}

{{ range $key, $value := .variables }}{{ $key }}={{ $value }}
{{ end }}