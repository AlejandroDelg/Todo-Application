package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func sendTodos(w http.ResponseWriter) {

	todos, err := getAllTodos()
	if err != nil {
		fmt.Println("Could not get all todos from db", err)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err = tmpl.Execute(w, todos)
	// err = tmpl.ExecuteTemplate(w, "Todos", todos)
	if err != nil {
		fmt.Println("Could not execute template", err)
	}
}



func index(w http.ResponseWriter, r * http.Request){
	sendTodos(w)
}

func isDoneTodo(w http.ResponseWriter, r * http.Request){
}

func deleteTodo(w http.ResponseWriter, r *http.Request){
	
}
func addTodo(w http.ResponseWriter, r *http.Request){
	
}
func setupRoutes(){
	muxRouter := mux.NewRouter()
	
	muxRouter.HandleFunc("/", index)
	muxRouter.HandleFunc("/todo{Id}", isDoneTodo).Methods("PUT")
	muxRouter.HandleFunc("/todo{Id}", deleteTodo).Methods("DELETE")
	muxRouter.HandleFunc("/addTodo", addTodo).Methods("POST")
	
	log.Fatal(http.ListenAndServe(":8080", muxRouter))
	
}