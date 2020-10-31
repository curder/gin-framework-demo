## gin重定向

- `Redirect` 路由重定向
- `HandleContext` 通过gin的上下文`c.Request.URL.Path`修改请求的路径，并调用`HandleContext`方法

```
go run main.go
```

打开本地浏览器访问：`http://localhost:8080` 将会被跳转到 `https://github.com/curder`

本地浏览器访问：`http://localhost:8080/a` 将会把 `http://localhost:8080/b` 地址的内容展示出来
                                                                               


