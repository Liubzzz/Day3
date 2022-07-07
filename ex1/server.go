package main

import (
	"ex1/db"
	"ex1/graph"
	"ex1/graph/generated"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	db.InitDatabase()
	router := chi.NewRouter()
	router.Handle("/query", handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
