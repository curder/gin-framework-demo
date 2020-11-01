## gorm连接MySQL基本示例


- 创建数据库`gin_framework_demo`

```
create database gin_framework_demo;
```

- 创建数据库表

```
db.AutoMigrate(&User{}) // 自动迁移
```

- 查询

```
var u User

db.First(&u)
```

- 更新

```                                  
var u User
db.Model(&u).Update("hobby", "music")
```

- 删除

```
var u User
db.Model(&u).Update("hobby", "music")
```


[GORM Docs](https://gorm.cn/docs/index.html)
