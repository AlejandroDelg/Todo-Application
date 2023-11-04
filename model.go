package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Todo struct {
	Id     int64
	Name   string
	IsDone bool
}
var db *sql.DB

var todos []Todo

func Setup() {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=root" +
		" dbname=todo sslmode=disable")
	//db, err = sql.Open("postgres", "host=localhost port=5433 user=root password=root dbname=todo sslmode=disable")

	if err != nil {
		log.Fatal("ERROR: ", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("ERROR en el ping: ", err)
	}

	log.Println("Connected to database")

	_, err = getAllTodos()
	if err != nil {
		return
	}

	todo, err := GetTodo(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("todo: ", todo)
}
func getAllTodos() ([]Todo, error){

	rows, err := db.Query("SELECT * FROM todos")
	if err != nil{
		log.Fatal("Error in selecting the query. Error: ", err)
	}

	var id int64
	var name string
	var isDone bool
	todos = []Todo{}
	
	for rows.Next(){
		err := rows.Scan(&id, &name, &isDone)
		if err != nil {
			return nil, err
		}
		actualTodo := Todo{
			Id:     id,
			Name:   name,
			IsDone: isDone,
		}
		todos = append(todos, actualTodo)
	}
	fmt.Println("<-------------TODOS------------->")
	for _, t := range todos {
		fmt.Println(t)
	}
	fmt.Println(" ")
	return todos, nil
}


func CreateTodo(name string)error{
	query := `Insert into todos(Id, Name, isdone) values ($1, $2, $3)`

	_, err := db.Exec(query, 2, name, false)
	if err != nil {
		return err
	}

	fmt.Println("Created todo")

	return nil

}

func DeleteTodo(id int64) error{
	query := `DELETE from todos WHERE Id = $1`

	_, err := db.Exec(query, id)

	if err != nil{
		fmt.Println("Error in deleting todo: " , err)
	}

	fmt.Println("todo of Id : " , id , " deleted")
	return nil

}

func GetTodo(id int64) (Todo, error){
	query := `Select * from todos where Id = $1`

	rows, err := db.Query(query, id)
	if err != nil {
		return Todo{}, err
	}
	var idGet int64
	var name string
	var isDone bool
	for rows.Next(){
		err := rows.Scan(&idGet, &name, &isDone)
		if err != nil {
			return Todo{}, err
		}
	}

	actualTodo := Todo{
		Id:     idGet,
		Name:   name,
		IsDone: isDone,
	}
	return actualTodo, nil
}

func SetDone(id int64) error{

	todo, err := GetTodo(id)
	if err != nil{
		return err
	}
	query := `update todos set isdone $1 where Id = $2`

	_, err = db.Exec(query,!todo.IsDone, id)
	if err != nil {
		return err
	}
	return nil
}
