package data

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var (
	DB *gorm.DB
)

//连接数据库函数
func Initmysql() (err error) {
	dsn := "root:ANTONG7230800@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	//if判断
	if err != nil {
		fmt.Println("打开数据库失败")
		return
	}
	//if之后的判断
	err = DB.DB().Ping()
	return
}
func Close(){
	DB.Close()
}