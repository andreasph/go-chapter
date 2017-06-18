package entity

type GoogleAutocomplete struct {
	Predictions []struct {
		Description string `json:"description"`
		PlaceId string `json:"place_id"`
	}
}
