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
	"github.com/go-chi/chi/middleware"
)

func main() {
	/* create router with chi */
	var router = new(chi.Mux)
	router = chi.NewRouter()

	/* sugguss middleware stack from chi */
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	/* use auth middleware */
	router.Use(auth.Middleware)

	/* init database */
	database.InitDB()

	/* add rest api register and login */
	router.Route("/api", func(r chi.Router) {
		r.Post("/register", auth.Register)
		r.Post("/login", auth.Login)
	})

	/* exec graphq schema and create server handler */
	executableSchema := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})
	server := handler.NewDefaultServer(executableSchema)

	/* router handler */
	router.Handle("/", playground.Handler("Starwars", "/query"))
	router.Handle("/query", server)

	/* logging on start and listen */
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.PORT)
	err := http.ListenAndServe(":"+config.PORT, router)
	if err != nil {
		panic(err)
	}
}
