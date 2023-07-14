package model

import (
	"fmt"
	"log"
)

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CreatePost(Firstname, Lastname string) error {
	insertQ, err := con.Query("INSERT INTO table1 VALUES(?,?)", Firstname, Lastname)
	errCheck(err)
	fmt.Println("dta done")
	insertQ.Close()
	return nil

}
