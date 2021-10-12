package client

import (
	"github.com/inkitio/gosdk/backend"
	"github.com/inkitio/gosdk/render"
)

type Client struct {
	RenderApi *render.RenderClient
}

func NewClient(apiKey string) *Client {

	backend := backend.NewBackend(apiKey)
	renderApi := render.NewRenderClient{backend: backend}

	return &Client{
		RenderApi: renderApi,
	}
}
