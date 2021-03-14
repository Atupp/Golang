package models

import (
	"demo/data"
)

//定义模板
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`  //代办事项内容
	Status bool   `json:"status"` //状态,完成/未完成
}

//todo增删改查
func Creatatodo(todo *Todo) (err error) {
	err = data.DB.Create(&todo).Error
	return
}
func Getalltodolist() (todolist []*Todo, err error) {
	if err = data.DB.Find(&todolist).Error; err != nil {
		return nil, err
	}
return
}

func GetAtodo(id string) (todo *Todo, err error) {
	if err = data.DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}
func UpdataAtodo(todo *Todo)(err error){
	err=data.DB.Save(todo).Error
	return
}
func DeleteATodo(id string)(err error){
	err=data.DB.Where("id=?",id).Delete(&Todo{}).Error
	return
}