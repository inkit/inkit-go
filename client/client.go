package client

import (
	"net/http"
	"time"
)

const (
	BaseURLV1 = "https://api.inkit.com/v1"
)

type Client struct {
	BaseUrl    string
	apiKey     string
	httpClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		BaseUrl: BaseURLV1,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
