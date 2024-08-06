package main

import (
	"api/controllers"
	"api/db"
	"api/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
  postgresDB := db.Connect()
  db.AutoMigrate(postgresDB)
  router := mux.NewRouter()
  router.HandleFunc("/api/users/create", controllers.CreateUser(postgresDB)).Methods("POST")
  router.HandleFunc("/api/users", controllers.GetUsers(postgresDB)).Methods("GET")
  router.HandleFunc("/api/users/{id}", controllers.GetUser(postgresDB)).Methods("GET")
  router.HandleFunc("/api/users/update/{id}", controllers.UpdateUser(postgresDB)).Methods("PUT")
  router.HandleFunc("/api/users/delete/{id}", controllers.DeleteUser(postgresDB)).Methods("DELETE")
  router.HandleFunc("/login", controllers.Login(postgresDB)).Methods("POST")
  enhancedRouter := middleware.EnableCors(middleware.JsonContentType(router))

  http.ListenAndServe(":8080", enhancedRouter)
}
