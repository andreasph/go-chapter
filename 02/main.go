package main

import (
	"github.com/gorilla/mux"
	"go-chapter/02/homework"
	"go-chapter/02/homework/handler"
	"go-chapter/02/google"
	"log"
	"net/http"
)

func main() {
	// read flags as a config
	config, err := homework.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	googleClient := google.NewClient(config.GoogleKey)

	r := mux.NewRouter()
	r.NotFoundHandler = handler.NewNotFound()

	// GET /cities-suggestions
	r.Methods(http.MethodGet).
		Path("/cities-suggestions").
		Handler(handler.NewCitiesSuggestions(googleClient))

	// POST /cities-info
	r.Methods(http.MethodPost).
		Path("/cities-info").
		Handler(handler.NewCitiesInfo(googleClient))

	// run HTTP server
	log.Fatalln(http.ListenAndServe(config.HTTPListen, handler.NewRecovery(r)))
}