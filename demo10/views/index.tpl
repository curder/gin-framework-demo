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
<h1>Hello {{ .name }}</h1>

<p>默认输出被转移后的字符串 {{ .unsafe }}</p>

<p>解析html输出 {{ .unsafe | safe }}</p>
</body>
</html>
