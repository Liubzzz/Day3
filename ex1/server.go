package main

import (
	"ex1/config"
	"ex1/db"
	"ex1/graph"
	"ex1/graph/generated"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {

	godotenv.Load("app.env")
	config.Initconfig()
	db.InitDatabase()
	router := chi.NewRouter()
	router.Handle("/"+config.GetConfig().APPVersion+"/query", handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})))
	router.Handle("/", playground.Handler("GraphQL playgroud", "/query"))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	log.Fatal(http.ListenAndServe(":"+config.GetConfig().Port, router))
}
