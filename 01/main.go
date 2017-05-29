package main

import "flag"
import "os"
import "go-chapter/01/homework"

func main() {
	var apiKey string
	var outputPath string
	var placeIds string

	flag.StringVar(&apiKey, "api-key", "", "Google Places API key")
	flag.StringVar(&outputPath, "output-path", "", "Path of the output file")
	flag.StringVar(&placeIds, "place-ids", "", "Comma separated list of Google Places IDS")

	flag.Parse()

	if len(apiKey) == 0 || len(outputPath) == 0 || len(placeIds) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	homework.Run(apiKey, outputPath, placeIds)
}
