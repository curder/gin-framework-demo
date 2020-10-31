<!doctype html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>{{ block "title" .}} 基础模版 {{ end }}</title>
</head>
<body>
{{/*定义一个自定义的区块，定义时这里的点"."必须加上，否则子模板将获取不到数据*/}}
{{ block "content" .}}这里是默认内容{{ end }}
</body>
</html>
