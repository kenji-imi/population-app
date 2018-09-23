{{- $short := (shortname .Name "err" "res" "sqlstr" "db" "XOLog") -}}
{{- $table := (.Table.TableName) -}}
class {{ .Name }} {
{{- range .Fields }}
	{{ if and (eq .Name "ID") (eq $table "population") -}}
		{{ .Name }} int32_AutoIncrement
	{{- else if eq .Name "Test" -}}
		{{ .Name }} TestType
	{{- else -}}
		{{ .Name }} {{ retype .Type }}
	{{- end }}
{{- end }}
}
