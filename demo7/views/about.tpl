{{ template "master.tpl" .}}

{{ define "title" }}这里是{{ . }}页面的自定义标题{{ end }}

{{ define "content" }}
    这里是 {{ . }} 页面的 自定义内容区域
{{ end }}
