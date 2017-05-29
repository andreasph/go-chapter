package entity

type Input struct {
	HtmlAttributions []string `json:"html_attributions"`
	Result           struct {
		AddressComponents []struct {
			LongName  string   `json:"long_name"`
			ShortName string   `json:"short_name"`
			Types     []string `json:"types"`
		} `json:"address_components"`
		AdrAddress       string `json:"adr_address"`
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			ViewPort struct {
				NorthEast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				SouthWest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"viewport"`
		} `json:"geometry"`
		Icon   string `json:"icon"`
		Id     string `json:"id"`
		Name   string `json:"name"`
		Photos []struct {
			Height           int      `json:"height"`
			HtmlAttributions []string `json:"html_attributions"`
			PhotoReference   string   `json:"photo_reference"`
			Width            int      `json:"width"`
		} `json:"photos"`
		PlaceId   string   `json:"place_id"`
		Reference string   `json:"reference"`
		Scope     string   `json:"scope"`
		Types     []string `json:"types"`
		Url       string   `json:"url"`
		UtcOffset int      `json:"utc_offset"`
		Vicinity  string   `json:"vicinity"`
	} `json:"result"`
	Status string `json:"status"`
}
