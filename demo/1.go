package main

import (
	"demo/data"
	"demo/models"
	"demo/routers"
	"fmt"

)



func main() {
	//创建数据库
	//SQL:creat database bubble;
	//连接数据据库
	err := data.Initmysql()
	if err != nil {
		fmt.Println("连接失败")
	}
	defer data.Close()
	//模型绑定
	data.DB.AutoMigrate(&models.Todo{})
	//接收用户请求
r:=routers.SetupRouter()
	r.Run(":8080")
}
