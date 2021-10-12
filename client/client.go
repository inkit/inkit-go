package client

import (
	"github.com/inkitio/gosdk/backend"
	"github.com/inkitio/gosdk/render"
)

type Client struct {
	Render *render.RenderClient
}

func NewClient(apiKey string) *Client {
	return &Client{
		Render: render.NewRenderClient(backend.NewBackend(apiKey)),
	}
}
