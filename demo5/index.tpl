<!doctype html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>自定义模版函数</title>
</head>
<body>

<h1>内置函数求长度</h1>
<p>String "{{ . }}" length: {{ len . }}</p>

<hr>

<h1>自定义模版函数</h1>

<p>{{ customFunc . }}</p>
</body>
</html>
