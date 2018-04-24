package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yaroslav-hnatyuk/stolen-democracy/infrastructure"
	"github.com/yaroslav-hnatyuk/stolen-democracy/interfaces/controllers"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(172.18.0.2:3306)/test-db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// we have two fields in table id and name so we need two placeholders or
	// else we will get an error
	stmtIns, err := db.Prepare("INSERT INTO names VALUES( ?, ? )")
	if err != nil {
		panic(err)
	}

	// Close the statement when we leave main() / the program terminates
	defer stmtIns.Close()

	// our id field auto increments so we don't need to pass actual value for it.
	_, err = stmtIns.Exec(nil, "John")
	if err != nil {
		panic(err)
	}

	// Prepare statement for reading data
	rows, err := db.Query("SELECT id, name FROM names")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var id int
	var name string

	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Printf("%d : %s \n", id, name)
	}

	controller := controllers.UserController{
		Logger: infrastructure.Logger{},
	}
	controller.GetUser(999)
}
