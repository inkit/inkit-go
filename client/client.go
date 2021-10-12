package client

import (
	backend "github.com/inkitio/gosdk/backend"
	render "github.com/inkitio/gosdk/render"
)

type Client struct {
	RenderApi *render.RenderClient
}

func NewClient(apiKey string) *Client {
	return &Client{
		RenderApi: render.NewRenderClient{backend: backend.NewBackend(apiKey)},
	}
}
