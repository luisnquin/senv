{{if .write_source_marker}}#_{{ .source_name }}_#

{{end}}{{ range $index, $variables := .grouped_variables }}{{ range $key, $value := $variables }}{{if $.use_export}}export {{end}}{{ $key }}="{{ $value }}"
{{end}}{{end}}