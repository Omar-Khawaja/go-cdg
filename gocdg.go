package gocdg

import (
	"errors"
	"net/http"
	"os"
	"time"
)

const (
	BaseURLv3              = "https://api.congress.gov/v3"
	ErrAPIKeyNotConfigured = "CDG_API_KEY is not set as environment variable"
)

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

func NewClient() (*Client, error) {
	cdgAPIKey := os.Getenv("CDG_API_KEY")
	if cdgAPIKey == "" {
		return nil, errors.New(ErrAPIKeyNotConfigured)
	}
	return &Client{
		BaseURL: BaseURLv3,
		apiKey:  cdgAPIKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}, nil
}
