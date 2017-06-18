package handler

import (
	"go-chapter/02/homework/request"
	"go-chapter/02/homework/response"
	"log"
	"net/http"
	"go-chapter/02/google"
	"go-chapter/02/entity"
)

type (
	citiesInfoHandler struct {
		googleClient *google.GoogleClient
	}
)

func NewCitiesInfo(googleClient *google.GoogleClient) http.Handler {
	return &citiesInfoHandler{googleClient}
}

func (h *citiesInfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// read POST request and handle an error
	var in []string
	if err := request.ParseBody(r.Body, &in); err != nil {
		response.WriteErrorString(w, http.StatusBadRequest, err.Error())
		return
	}

	// nothing to do
	if len(in) == 0 {
		response.WriteErrorString(w, http.StatusBadRequest, "You must specify at least one place id")
		return
	}

	var payload []entity.Place

	for _, id := range in {
		// build request
		// perform and handle an error
		place, err := h.googleClient.GetPlace(id)
		if err != nil {
			log.Printf("Error on get place %q details: %v\n", id, err)
			continue
		}

		payload = append(payload, place)
	}

	response.WriteSuccess(w, payload)
}