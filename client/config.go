package client

import (
	"errors"
	"os"
)

const endpointDefault = "https://places.googleapis.com/v1/places:searchNearby"

type Config struct {
	Endpoint string
	ApiToken string
}

func NewConfig() (*Config, error) {
	apiToken := os.Getenv("PLACES_CLIENT_API_TOKEN")
	endpoint := os.Getenv("PLACES_CLIENT_ENDPOINT")

	if endpoint == "" {
		endpoint = endpointDefault
	}

	if apiToken == "" {
		return nil, errors.New("environment variable PLACES_CLIENT_API_TOKEN is not set")
	}

	return &Config{
		Endpoint: endpoint,
		ApiToken: apiToken,
	}, nil
}
