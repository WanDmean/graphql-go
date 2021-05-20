package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/WanDmean/graphql-go/graph"
	"github.com/WanDmean/graphql-go/graph/generated"
	"github.com/WanDmean/graphql-go/src/auth"
	"github.com/WanDmean/graphql-go/src/config"
	"github.com/WanDmean/graphql-go/src/database"
	"github.com/go-chi/chi"
)

func main() {
	port := config.PORT

	/* create router with chi */
	router := chi.NewRouter()

	/* use auth middleware */
	router.Use(auth.Middleware())

	/* init database */
	database.InitDB()

	/* exec graphq schema and create server handler */
	executableSchema := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})
	server := handler.NewDefaultServer(executableSchema)

	/* router handler */
	router.Handle("/", playground.Handler("Starwars", "/query"))
	router.Handle("/query", server)

	/* logging on start and listen */
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
