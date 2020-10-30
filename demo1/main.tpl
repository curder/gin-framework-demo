<!doctype html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
<h1>Hello Golang</h1>
<h2 style="color: green">hello world</h2>
<button id="btn">Click</button>
<img id="img" src="https://www.google.com/logos/2020/halloween20/rc1/cta.png" alt="图片">
<script>
    window.onload = () => {
        const btn = document.getElementById("btn");
        const img = document.getElementById("img")

        btn.onclick = (e) => {
            img.src = 'https://lh3.googleusercontent.com/ogw/ADGmqu-zru3SnqyRXjCqFJvVDnO7Vs_Z8YB6nDHHjF3e=s32-c-mo';
        }
    }
</script>
</body>
</html>
