#_{{ .sourceName }}_#

{{ range $key, $value := .variables }}{{ $key }}={{ $value }}
{{ end }}