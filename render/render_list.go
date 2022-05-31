package render

import (
	"bytes"
	"encoding/json"

	"github.com/inkit/inkit-go/backend"
)

func (c *RenderClient) List(options *RenderListOptions) (*RendersList, error) {
	rendersList := &RendersList{}

	var backendOptions *backend.ListOptions

	if options != nil {
		backendOptions = &backend.ListOptions{
			Page:     options.Page,
			PageSize: options.PageSize,
		}
	}

	var resPayload []byte
	resPayload, err := c.Backend.ListRequest("/renders", backendOptions)

	if err != nil {
		return nil, err
	}

	if err = json.NewDecoder(bytes.NewBuffer(resPayload)).Decode(rendersList); err == nil {
		return rendersList, nil
	} else {
		return nil, err
	}

}
