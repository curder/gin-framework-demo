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

{{/*从结构体中获取数据*/}}
<h1>Render Struct Data: </h1>
<p>Name: {{ .u1.Name }}</p>
<p>Age: {{ .u1.Age }}</p>
<p>Gender: {{ .u1.Gender }}</p>

<hr>
{{/*从map中获取变量的值*/}}
<h1>Render Map data:</h1>
<p>Name: {{ .m1.name }}</p>
<p>Age: {{ .m1.age }}</p>
<p>Gender: {{ .m1.gender }}</p>

<hr>

{{/*自定义变量*/}}
<h1>Template Variables:</h1>
{{ $user := .m1 }} {{/*变量的定义*/}}
<p>Name: {{- $user.name -}} {{/*移除左右空格*/}}</p>
<p>Age: {{ $user.age }} {{/*变量的使用*/}}</p>
<p>Gender: {{ $user.gender }}</p>


<hr>

<h1>If Conditions: </h1>
{{ $age := 18 }}
{{ if le $age 18 }}
    Current age value is {{ $age }}, study now.
{{ else }}
    Current age value is {{ $age }}, work hard.
{{ end }}

<hr>

<h1>Range: </h1>
{{ range $index, $hobby := .m1.hobbyList }}
    <p>index: {{ $index }} - {{ $hobby }}</p>
{{ end }}

<hr>

<h1>With: </h1>
{{ with .u1 }}
    <p>Name: {{ .Name }}</p>
    <p>Age: {{ .Age }}</p>
    <p>Gender: {{ .Gender }}</p>
{{ end }}

</body>
</html>
