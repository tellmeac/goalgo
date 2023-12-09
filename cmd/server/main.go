package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tellmeac/goalgo/internal/app"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	service := app.New(&app.Config{DatabaseConn: os.Getenv("POSTGRES_URI")})
	delivery := app.NewDelivery(service)

	router := mux.NewRouter()
	router.Path("/api/updates").Methods(http.MethodGet).HandlerFunc(delivery.GetUpdates)
	router.Path("/api/chart").Methods(http.MethodGet).HandlerFunc(delivery.GetChart)

	log.Print("Starting server")

	go log.Fatal(http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodOptions,
		}),
	)(router)))

	webHandler := http.FileServer(http.Dir("./static"))

	log.Fatal(http.ListenAndServe(":8000", webHandler))
}
