package controller

import (
	"demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)
/*
其实现实是这样的
url-->controller-->logic-->model
请求-->控制器-->业务逻辑-->模型层的增删改查
 */
func Indexcontroller(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Creattodo(c *gin.Context) {
	//输入内容,返回这里
	var todo models.Todo
	c.BindJSON(&todo)
	//内容放到数据库
err:=models.Creatatodo(&todo)
if err!=nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
	//返回响应
}

func Gettodolist(c *gin.Context) {
	todolist,err:=models.Getalltodolist()
	if err!= nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}

func Updateatodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效ID"})
	}
	todo,err:=models.GetAtodo(id)
	if err!=nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err = models.UpdataAtodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func Deleteatodo(c *gin.Context) {
	//拿到ID参数
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效ID"})
	}
	if err := models.DeleteATodo(id);err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}