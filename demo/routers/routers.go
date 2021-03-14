package routers

import (
	"demo/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter()*gin.Engine{
	r := gin.Default()
	//加载静态文件
	r.Static("/static", "static")
	//加载模板
	r.LoadHTMLGlob("templates/*")
	r.GET("/upp", controller.Indexcontroller)
	//V1
	v1Group := r.Group("v1")

	//所有待办事项

	//查看
	v1Group.GET("/todo", controller.Gettodolist)
	//添加
	v1Group.POST("/todo",controller.Creattodo)

	//单一待办事项

	//修改
	v1Group.PUT("/todo/:id", controller.Updateatodo)
	//删除
	v1Group.DELETE("/todo/:id", controller.Deleteatodo)
	return r
}
