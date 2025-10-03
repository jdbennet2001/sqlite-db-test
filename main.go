package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "./comics.db")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to: ", db)
}
