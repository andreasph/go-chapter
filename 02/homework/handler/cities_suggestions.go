package handler

import (
	"go-chapter/02/homework/response"
	"log"
	"net/http"
	"go-chapter/02/google"
	"go-chapter/02/entity"
)

type (
	citiesSuggestionsHandler struct {
		googleClient *google.GoogleClient
	}

	citiesSuggestionsItem struct {
		PlaceID string `json:"place_id"`
		Name    string `json:"name"`
	}
)

func NewCitiesSuggestions(googleClient *google.GoogleClient) http.Handler {
	return &citiesSuggestionsHandler{googleClient}
}

func (h *citiesSuggestionsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// validate query param
	query := r.URL.Query().Get("q")
	if query == "" {
		response.WriteErrorString(w, http.StatusBadRequest, "Query parameter is missing")
		return
	}

	// perform and handle an error
	res, err := h.googleClient.PlaceAutocomplete(query)
	if err != nil {
		response.WriteErrorString(w, http.StatusInternalServerError, "Unexpected error on fetching cities")
		log.Printf("CitiesSuggestions: %v\n", err)
		return
	}

	// gather necessary information
	payload := make([]entity.Suggestion, len(res.Predictions))
	for i, p := range res.Predictions {
		payload[i] = entity.Suggestion{
			PlaceId: p.PlaceId,
			Name:    p.Description,
		}
	}

	response.WriteSuccess(w, payload)
}
