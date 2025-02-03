#_{{ .source_name }}_#

{{ range $index, $variables := .grouped_variables }}{{ range $key, $value := $variables }}{{if $.use_export}}export {{end}}{{ $key }}="{{ $value }}"
{{end}}{{end}}