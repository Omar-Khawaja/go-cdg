package gocdg

import (
	"encoding/json"
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

type service struct {
	*Client
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

func (c *Client) newRequest(req *http.Request, v interface{}) error {
	q := req.URL.Query()
	q.Add("api_key", c.apiKey)
	req.URL.RawQuery = q.Encode()

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}
