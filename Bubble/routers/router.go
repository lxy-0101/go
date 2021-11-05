package routers

import (
	"Bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static","static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/",controller.IndexHandler)
	v1group := r.Group("v1")
	{
		v1group.GET("/todo",controller.GetTodoList)
		v1group.POST("/todo", controller.CreateTodo)
		v1group.PUT("/todo/:id",controller.UpdateATodo)
		v1group.DELETE("/todo/:id",controller.DeleteATodo)
	}
	return r
}
