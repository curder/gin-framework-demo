## gin获取post请求主体参数

- `PostForm` 获取Post传递主体的值，不传递则为空
- `DefaultPostForm` 获取Post请求主体传递的key的值，如果没有key的值，则值为默认值
- `GetPostForm` 获取Post请求主体传递的key的值，返回key的值和是否传递


```
go run main.go
```           

打开本地浏览器访问：`http://127.0.0.1:8080`


