package client

import (
	"github.com/inkit/inkit-go/backend"
	"github.com/inkit/inkit-go/render"
)

type Client struct {
	Render *render.RenderClient
}

func NewClient(apiKey string) *Client {
	return &Client{
		Render: render.NewRenderClient(backend.NewBackend(apiKey)),
	}
}
