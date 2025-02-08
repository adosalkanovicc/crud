package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adosalkanovicc/go_crud/config"
	"github.com/adosalkanovicc/go_crud/handler"
)

func main() {
	db := config.InitDB()
	defer db.Close()

	router := handler.SetupRoutes(db)

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
