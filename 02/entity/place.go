package entity

type Place struct {
	LongName    string `json:"long_name"`
	ShortName   string `json:"short_name"`
	Coordinates struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"coordinates"`
	FormattedAddress string  `json:"formatted_address"`
	PlaceId          string  `json:"place_id"`
	Rating           float64 `json:"rating"`
	PhotoReference   string  `json:"photo_reference"`
}

func (output *Place) FromInput(input GooglePlaces) {
	if (len(input.Result.AddressComponents) > 0) {
		output.LongName = input.Result.AddressComponents[0].LongName
		output.ShortName = input.Result.AddressComponents[0].ShortName
	}
	output.Coordinates.Lat = input.Result.Geometry.Location.Lat
	output.Coordinates.Lng = input.Result.Geometry.Location.Lng
	output.FormattedAddress = input.Result.FormattedAddress

	if (len(input.Result.Photos) > 0) {
		output.PhotoReference = input.Result.Photos[0].PhotoReference
	}

	output.PlaceId = input.Result.PlaceId
}
