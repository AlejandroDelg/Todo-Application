package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type todo struct {
	id 	int64
	name 	string
	isDone	bool
}
var db *sql.DB

func Setup() {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=root" +
		" dbname=todo sslmode=disable")
	//db, err = sql.Open("postgres", "host=localhost port=5433 user=root password=root dbname=todo sslmode=disable")

	todos := [] todo{}
	if err != nil {
		log.Fatal("ERROR: ", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("ERROR en el ping: ", err)
	}

	log.Println("Connected to database")

	rows, err := db.Query("SELECT * FROM todos")
	if err != nil{
		log.Fatal("Error in selecting the query. Error: ", err)
	}

	var id int64
	var name string
	var isDone bool
	
	for rows.Next(){
		err := rows.Scan(&id, &name, &isDone)
		if err != nil{
			fmt.Println("Error: ", err)
		}

		actualTodo := todo{
			id:     id,
			name:   name,
			isDone: isDone,
			}

			todos = append(todos,actualTodo)
	}

	for _, t := range todos {
		fmt.Println(t)
	}
	

	defer db.Close()
}
