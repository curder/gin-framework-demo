## gin获取queryString

- `Query` 获取地址栏传递的key的值，不传递则为空
- `DefaultQuery` 如果没有key的值，则值curder值
- `GetQuery` 获取地址栏传递的key的值，返回key的值和是否传递

```
go run main.go
```           

打开本地浏览器访问：`http://127.0.0.1:8080`或者 `http://127.0.0.1:8080/?search=abc`

