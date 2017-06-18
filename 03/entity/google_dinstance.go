package entity

type GoogleDistance struct {
	Rows []struct {
		Elements []struct {
			Distance  struct {
				Value float64   `json:"value"`
			}   `json:"distance"`
			Status string `json:"status"`
		} `json:"elements"`
	} `json:"rows"`
	Status string `json:"status"`
}
