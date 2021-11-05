package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	DB *gorm.DB
)
func InitMySql() (err error) {
	//连接数据库
	dsn := "root:123456@tcp(localhost)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB,err = gorm.Open("mysql",dsn)
	if err != nil{
		return
	}
	return DB.DB().Ping()
}
func Close()  {
	DB.Close()
}
