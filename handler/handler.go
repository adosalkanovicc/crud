package handler

import (
	"database/sql"

	"github.com/adosalkanovicc/go_crud/controller"
	"github.com/adosalkanovicc/go_crud/repository"
	"github.com/adosalkanovicc/go_crud/service"
	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	return router
}
