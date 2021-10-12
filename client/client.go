package client

import (
	"backend"
	"render"
)

type Client struct {
	RenderApi *render.Render
}

func NewClient(apiKey string) *Client {
	return &Client{
		RenderApi: render.NewRenderClient{backend: backend.NewBackend(apiKey)},
	}
}
