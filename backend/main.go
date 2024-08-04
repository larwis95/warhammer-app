package main

import (
	"api/graph"
	"fmt"
	"log"
	"net/http"
	"os"

	"api/models"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
  db := connectDB()
  fmt.Println(db.Name())
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}


func connectDB() *gorm.DB {
  dbURL := os.Getenv("DATABASE_URL")
  db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
  db.AutoMigrate(&models.User{})
  db.AutoMigrate(&models.Allegiance{})
  db.AutoMigrate(&models.GrandAlliance{})
  db.AutoMigrate(&models.Unit{})

  return db
}
