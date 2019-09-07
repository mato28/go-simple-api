package models

import (
	u "awesomeProject/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Name string `json:"name"`
	Description string `json:"description"`
	UserId uint `json:"user_id"` //The user that this contact belongs to
}

func (todo *Todo) Validate() (map[string]interface{}, bool) {

	if todo.Name == "" {
		return u.Message(false, "Contact name should be on the payload"), false
	}

	if todo.Description == "" {
		return u.Message(false, "Phone number should be on the payload"), false
	}

	if todo.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (todo *Todo) CreateTodo() map[string] interface{} {
	if resp, ok := todo.Validate(); !ok {
		return resp
	}

	GetDB().Create(todo)

	resp := u.Message(true, "success")
	resp["todo"] = todo
	return resp
}

func GetAllTodos() []Todo {
	todos := []Todo {}
	err := GetDB().Find(&todos).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return todos
}

func GetTodo(id int) *Todo {
	todo := &Todo{}
	if err := GetDB().First(todo, id).Error; err != nil {
		fmt.Println(err)
		return nil
	}

	return todo
}

func (todo *Todo) Update(id int) *Todo {
	if _, ok := todo.Validate(); !ok {
		return nil
	}
	currentTodo := &Todo{}
	err := GetDB().First(currentTodo, id).Error
	if err != nil {
		fmt.Print(err)
		return nil
	}
	currentTodo.Name = todo.Name
	currentTodo.Description = todo.Description
	GetDB().Save(currentTodo)
	return currentTodo
}

func DeleteTodo(id int) bool {
	todo := &Todo{}
	if err := GetDB().First(todo, id).Error; err != nil {
		return false
	}
	GetDB().Delete(todo)
	return true
}