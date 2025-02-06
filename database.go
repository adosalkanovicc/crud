package main

import (
	"database/sql"
	"fmt"
	"log"
)

func InitDB() {
	var err error
	dsn := "go_user:password123@tcp(localhost:3306)/userdb"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Can't reach the database:", err)
	}
	fmt.Println("Connected to the database!")
}
