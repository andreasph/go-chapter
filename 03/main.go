package main

import (
	"go-chapter/03/homework"
	"go-chapter/03/google"
	"log"
	"fmt"
)

func main() {
	// read flags as a config
	config, err := homework.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	googleClient := google.NewClient(config.GoogleKey)

	res, err := googleClient.GetDistances(config.PlaceIds)

	fmt.Printf("distances: %v, error: %v", res, err)
}

