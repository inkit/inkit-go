package render

import (
	"backend"
	"bytes"
	"encoding/json"
)

func (c *RenderClient) ListRenders(options *backend.ListOptions) (*RendersList, error) {
	rendersList := &RendersList{}

	var resPayload []byte
	resPayload, err := c.Backend.ListRequest("/renders", options)

	if err != nil {
		return nil, err
	}

	if err = json.NewDecoder(bytes.NewBuffer(resPayload)).Decode(rendersList); err == nil {
		return rendersList, nil
	} else {
		return nil, err
	}

}
