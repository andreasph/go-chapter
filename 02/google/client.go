package google

import (
	j "encoding/json"
	"go-chapter/02/entity"
	"io/ioutil"
	"net/http"
	"time"
	"fmt"
)
type GoogleClient struct {
	apiKey string
	googlePlaceUrl string
}

func NewClient(apiKey string) *GoogleClient {
	return &GoogleClient{apiKey,"https://maps.googleapis.com/maps/api/place/%s/json?key=%s"}
}

func (client *GoogleClient) PlaceAutocomplete(query string) (entity.GoogleAutocomplete, error) {
	var res = entity.GoogleAutocomplete{}

	currentUrl := client.googlePlaceUrl + "&input=%s&types=(cities)&language=en_EN"
	currentUrl = fmt.Sprintf(currentUrl, "autocomplete", client.apiKey, query)
	json, err := downloadJson(currentUrl)
	if err != nil {
		return res, err
	}

	err = j.Unmarshal(json, &res)
	if err != nil {
		return res, err
	}

	return res, nil;
}

func (client *GoogleClient) GetPlace(placeId string) (entity.Info, error) {
	var res = entity.Info{}

	currentUrl := client.googlePlaceUrl + "&placeid=%s"
	currentUrl = fmt.Sprintf(currentUrl, "details", client.apiKey, placeId)
	json, err := downloadJson(currentUrl)
	if err != nil {
		return res, err
	}

	var input entity.GooglePlaces

	err = j.Unmarshal(json, &input)
	if err != nil {
		return res, err
	}

	res.FromInput(input)

	return res, nil;
}

func downloadJson(url string) ([]byte, error) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	response, err := netClient.Get(url)

	if err != nil {
		return []byte("{}"), err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte("{}"), err
	}

	return body, nil

}

