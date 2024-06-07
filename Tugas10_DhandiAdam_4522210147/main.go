package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Todo struct {
	ID        string `json:"id"`
	TitleId   string `json:"title_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos []Todo

func main() {
	router := mux.NewRouter()

	//Menambahkan handler untuk route "/todos"
	router.HandleFunc("/todos", GetTodos).Methods("GET")
	router.HandleFunc("/todos", CreateTodo).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

//Handler untuk mendapatkan semua todos

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var newTodo Todo
	json.NewDecoder(r.Body).Decode(&newTodo)
	newTodo.ID = uuid.New().String()
	todos = append(todos, newTodo)
	json.NewEncoder(w).Encode(newTodo)
}
