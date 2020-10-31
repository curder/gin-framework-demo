<!doctype html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Index</title>
</head>
<body>
{{/*外部模版*/}}
<h1>From ul.tpl file</h1>
{{ template "ul.tpl"}}
<hr>

{{/*当前模版定义的模版块*/}}
<h1>From Current file</h1>
{{ template "current.tpl"}}

<hr>

{{/*模版引擎数据*/}}
{{ . }}
</body>
</html>
{{ define "current.tpl"}}
    <ol>
        <li>Learn</li>
        <li>Eat</li>
        <li>Work</li>
    </ol>
{{ end }}
