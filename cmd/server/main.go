package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/tellmeac/goalgo/internal/app"
	"log"
	"net/http"
	"os"
)

func main() {
	service := app.New(&app.Config{DatabaseConn: os.Getenv("POSTGRES_URI")})
	delivery := app.NewDelivery(service)

	router := mux.NewRouter()
	router.Path("/api/updates").Methods(http.MethodGet).HandlerFunc(delivery.GetUpdates)
	router.Path("/api/chart").Methods(http.MethodGet).HandlerFunc(delivery.GetChart)

	log.Print("Starting server")

	go func() {
		server := &http.Server{
			Addr: ":8080",
			Handler: handlers.CORS(
				handlers.AllowedOrigins([]string{"*"}),
				handlers.AllowedMethods([]string{
					http.MethodGet,
					http.MethodHead,
					http.MethodPost,
					http.MethodPut,
					http.MethodOptions,
				}),
			)(router),
		}

		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	webPath := os.Getenv("WEB_PATH")
	if webPath == "" {
		log.Print("empty WEB_PATH")
		return // TODO: signals
	}

	log.Fatal(http.ListenAndServe(":8000", http.FileServer(http.Dir(webPath))))
}
