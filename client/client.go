package client

import (
	"github.com/inkitio/gosdk/backend"
	"github.com/inkitio/gosdk/render"
)

type Client struct {
	RenderApi *render.RenderClient
}

func NewClient(apiKey string) *Client {
	return &Client{
		RenderApi: render.NewRenderClient(backend.NewBackend(apiKey)),
	}
}
