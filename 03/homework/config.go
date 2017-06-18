package homework

import (
	"github.com/jessevdk/go-flags"
)

type Config struct {
	GoogleKey  string `long:"key" description:"Google Maps API Key" required:"true"`
	PlaceIds  string `long:"place_ids" description:"Comma separated list of Google Maps Place IDs" required:"true"`
}

// NewConfig return the instance of config structure or error if so
func NewConfig() (*Config, error) {
	var c Config

	if _, err := flags.NewParser(&c, flags.HelpFlag|flags.PassDoubleDash).Parse(); err != nil {
		return nil, err
	}

	return &c, nil
}