package google

import (
	j "encoding/json"
	"go-chapter/03/entity"
	"io/ioutil"
	"net/http"
	"time"
	"fmt"
	"strings"
	"errors"
	"math"
)

type GoogleClient struct {
	apiKey string
	googlePlaceUrl string
}

func NewClient(apiKey string) *GoogleClient {
	return &GoogleClient{apiKey,"https://maps.googleapis.com/maps/api/%s/json?key=%s"}
}

func (client *GoogleClient) GetDistances(inputIds string) (map[string]map[string]float64, error) {
	placeIds := strings.Split(inputIds, ",")
	for key, placeId := range placeIds {
		placeIds[key] = strings.Trim(placeId, " ")
	}

	placesInfo := make([]entity.Info, len(placeIds))
	for i, placeId := range placeIds {
		placeIds[i] = fmt.Sprintf("place_id:%s", placeId)
		placesInfo[i], _ = client.GetPlace(placeId)
	}

	query := strings.Join(placeIds, "|")


	res := make(map[string]map[string]float64)

	currentUrl := client.googlePlaceUrl + "&origins=%s&destinations=%s"
	currentUrl = fmt.Sprintf(currentUrl, "distancematrix", client.apiKey, query, query)
	json, err := downloadJson(currentUrl)

	if err != nil {
		return res, err
	}

	var input entity.GoogleDistance

	err = j.Unmarshal(json, &input)
	if err != nil {
		return res, err
	}

	if input.Status != "OK" {
		return res, errors.New("Error in getting google distances")
	}

	for i, row := range input.Rows {
		for ii, element := range row.Elements {
			if element.Status != "OK" {
				continue
			}

			if (i == ii) {
				continue
			}

			origin := placesInfo[i].ShortName
			destination := placesInfo[ii].ShortName

			if _, ok := res[destination][origin]; ok {
				continue
			}

			if _, ok := res[origin]; !ok {
				res[origin] = make(map[string]float64)
			}

			if _, ok := res[origin][destination]; !ok {

				res[origin][destination] = math.Floor(element.Distance.Value / 1000);
			}
		}
	}

	return res, nil;
}

func (client *GoogleClient) GetPlace(placeId string) (entity.Info, error) {
	var res = entity.Info{}

	currentUrl := client.googlePlaceUrl + "&placeid=%s"
	currentUrl = fmt.Sprintf(currentUrl, "place/details", client.apiKey, placeId)
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

