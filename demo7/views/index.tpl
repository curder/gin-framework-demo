{{ template "master.tpl" .}}

{{ define "title"}}首页 {{ . }} {{ end}}

{{ define "content"}}
    这里是首页的内容 {{ . }}
{{end}}
