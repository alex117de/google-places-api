package client

import "github.com/alex117de/google-places-api/logger"

const EndpointDefault = "https://places.googleapis.com/v1/places:searchNearby"

type Config struct {
	Endpoint string
	ApiToken string
	LogLevel logger.LogLevel
}
