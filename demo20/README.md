## gin路由和路由分组

- `NoRoute` 统一定义404的响应
- `Any` 同时可以注册多种请求方式的处理，可以通过switch-case语法来分别处理每个请求
- `Group` 对路由进行分组，一般将一些前缀一致的路由进行分组，方便统一管理
```
go run main.go
```

打开本地浏览器访问：`http://localhost:8080`
