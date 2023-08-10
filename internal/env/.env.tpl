#_{{ .sourceName }}_#

{{ range $key, $value := .variables }}{{if $.useExport}}export {{end}}{{ $key }}="{{ $value }}"
{{ end }}