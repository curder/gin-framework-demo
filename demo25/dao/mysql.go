package dao

import "github.com/jinzhu/gorm"

var DB *gorm.DB

// InitDB
func InitDB() (err error) {

	if DB, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:33060)/gin_framework_demo?charset=utf8mb4&parseTime=True&loc=Local"); err != nil {
		return
	}

	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}
