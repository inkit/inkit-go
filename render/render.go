package render

import (
	"github.com/inkit/inkit-go/backend"
)

type RenderClient struct {
	Backend *backend.Backend
}

func NewRenderClient(backend *backend.Backend) *RenderClient {
	return &RenderClient{
		Backend: backend,
	}
}
