package render

import (
	"bytes"
	"encoding/json"
	"github.com/inkitio/gosdk/client"
	"github.com/inkitio/gosdk/inkit"
)

func (c *client.Client) ListRenders(options *inkit.ListOptions) (*RendersList, error) {
	rendersList := &RendersList{}

	var resPayload []byte
	resPayload, err := c.ListRequest("/renders", options)

	if err != nil {
		return nil, err
	}

	if err = json.NewDecoder(bytes.NewBuffer(resPayload)).Decode(rendersList); err == nil {
		return rendersList, nil
	} else {
		return nil, err
	}

}
