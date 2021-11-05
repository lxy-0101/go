package main

import (
	"Bubble/dao"
	"Bubble/models"
	"Bubble/routers"
	"fmt"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "net/http"
)

func main() {
	err := dao.InitMySql()
	if err != nil{
		fmt.Println("init mysql failed",err)
	}
	defer dao.DB.Close()
	//建表
	dao.DB.AutoMigrate(&models.Todo{})
	//u1 := &Todo{1,"happy",true}
	//DB.Create(u1)
	//DB.Delete(u1)

	r := routers.SetUpRouter()
	r.Run()

}