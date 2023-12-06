package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tellmeac/goalgo/internal/app"
	"log"
	"net/http"
)

func main() {
	service := app.New(&app.Config{DatabaseConn: "./sqlite.db"})
	delivery := app.NewDelivery(service)

	router := mux.NewRouter()
	router.Path("/updates").Methods(http.MethodGet).HandlerFunc(delivery.GetUpdates)
	router.Path("/chart").Methods(http.MethodGet).HandlerFunc(delivery.GetChart)

	log.Fatal(http.ListenAndServe(":80", handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}),
	)(router)))
}
