package model

import (
	"database/sql"
	"fmt"
	"log"
)

type Todo struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:Dhanu123@tcp(127.0.0.1:3306)/prg")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("database connected")
	con = db

	return db
}
