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
  defer db.AutoMigrate(postgresDB)
  router := mux.NewRouter()
  authRoutes := router.Methods("POST", "GET", "PUT", "DELETE").Subrouter()
  router.HandleFunc("/api/users/create", controllers.CreateUser(postgresDB)).Methods("POST")
  authRoutes.HandleFunc("/api/users", controllers.GetUsers(postgresDB)).Methods("GET")
  authRoutes.HandleFunc("/api/users/{id}", controllers.GetUser(postgresDB)).Methods("GET")
  authRoutes.HandleFunc("/api/users/update/{id}", controllers.UpdateUser(postgresDB)).Methods("PUT")
  authRoutes.HandleFunc("/api/users/delete/{id}", controllers.DeleteUser(postgresDB)).Methods("DELETE")
  router.HandleFunc("/login", controllers.Login(postgresDB)).Methods("POST")
  router.HandleFunc("/api/units/compare", controllers.CompareUnits()).Methods("GET")
  authRoutes.Use(middleware.AuthMiddleware)
  enhancedRouter := middleware.EnableCors(middleware.JsonContentType(router))


  http.ListenAndServe(":8080", enhancedRouter)
}
