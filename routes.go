package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func sendTodos(w http.ResponseWriter, r * http.Request){
}


func index(w http.ResponseWriter, r * http.Request){

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Error in executing template")
		return 
	}
	
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
	muxRouter.HandleFunc("/todo{id}", isDoneTodo).Methods("PUT")
	muxRouter.HandleFunc("/todo{id}", deleteTodo).Methods("DELETE")
	muxRouter.HandleFunc("/todo/{id", addTodo).Methods("POST")
	
	log.Fatal(http.ListenAndServe(":8080", muxRouter))
	
}