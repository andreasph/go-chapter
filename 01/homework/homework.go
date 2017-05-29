package homework

import (
	j "encoding/json"
	"go-chapter/01/entity"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Run(apiKey string, outputPath string, inputIds string) {
	const googlePlaceUrl = "https://maps.googleapis.com/maps/api/place/details/json?placeid=%PLACE_ID%&key=%API_KEY%"

	placeIds := strings.Split(inputIds, ",")
	for key, placeId := range placeIds {
		placeIds[key] = strings.Trim(placeId, " ")
	}

	var outputs []entity.Output

	for _, placeId := range placeIds {
		currentUrl := strings.Replace(googlePlaceUrl, "%API_KEY%", apiKey, -1)
		currentUrl = strings.Replace(currentUrl, "%PLACE_ID%", placeId, -1)
		json, err := downloadJson(currentUrl)
		if err != nil {
			continue
		}

		var input entity.Input

		err = j.Unmarshal(json, &input)
		if err != nil {
			continue
		}

		var current = entity.Output{}
		current.FromInput(input)

		outputs = append(outputs, current)
	}

	output, _ := j.MarshalIndent(outputs, "", "	")
	ioutil.WriteFile(outputPath, output, 0644)
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
