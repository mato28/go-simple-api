package controllers

import (
	"awesomeProject/models"
	u "awesomeProject/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var CreateTodo = func(w http.ResponseWriter, r *http.Request) {

	todo := &models.Todo{}

	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := todo.CreateTodo()
	u.Respond(w, resp)
}

var GetTodos = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetAllTodos()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetTodo = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetTodo(id)
	if data != nil {
		resp := u.Message(true, "success")
		resp["data"] = data
		u.Respond(w, resp)
	} else {
		response := make(map[string] interface{})
		response = u.Message(false, "Not found")
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("Content-Type", "application/json")
		u.Respond(w, response)
	}
}

var UpdateTodo = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	todo := &models.Todo{}

	errDec := json.NewDecoder(r.Body).Decode(todo)
	if errDec != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	data := todo.Update(id)
	if data != nil {
		resp := u.Message(true, "success")
		resp["data"] = data
		u.Respond(w, resp)
	} else {
		response := make(map[string] interface{})
		response = u.Message(false, "Not found")
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("Content-Type", "application/json")
		u.Respond(w, response)
	}
}

var DeleteTodo =  func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}
	if models.DeleteTodo(id) {
		u.Respond(w, u.Message(true, "success"))
	} else {
		u.Respond(w, u.Message(false, "Not found"))
	}
}