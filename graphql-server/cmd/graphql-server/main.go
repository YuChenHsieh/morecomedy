package main

import (
	"graphql-server/internal/graph"
	"graphql-server/internal/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {
	mux := mux.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	// db.ClickhouseInit()
	// db.MinioInit()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	mux.Handle("/query", srv)
	mux.HandleFunc("/playground", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/", mux)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(mux)))
}
