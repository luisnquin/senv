# source: {{ .sourceName }}

{{ range $key, $value := .values }}{{ $key }}={{ $value }}
{{ end }}