package main

import (
	"log"
	"my_gql_server/app/database"
	"my_gql_server/app/repository"
	"my_gql_server/graph"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"

	"gorm.io/gorm"
)

const defaultPort = "8080"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	godotenv.Load()
	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := database.NewConnection(config)
	if err != nil {
		panic(err)
	}
	database.Migrate(db)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	repo := repository.NewBookService(db)
	panRepo := repository.NewPanService(db)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{BookRepository: repo, PanRepository: panRepo}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
