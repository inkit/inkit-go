package render

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetRender(id string) (*Render, error) {
	render := &Render{}

	var resPayload []byte
	resPayload, err := c.GetRequest(fmt.Sprintf("/render/%s", id))

	if err != nil {
		return nil, err
	}

	if err = json.NewDecoder(bytes.NewBuffer(resPayload)).Decode(render); err == nil {
		return render, nil
	} else {
		return nil, err
	}
}

func (c *Client) GetRenderPdf(id string) (*[]byte, error) {
	var resPayload []byte

	resPayload, err := c.GetRequest(fmt.Sprintf("/render/%s/pdf", id))

	if err != nil {
		return nil, err
	}

	return &resPayload, nil
}

func (c *Client) GetRenderHtml(id string) (*[]byte, error) {
	var resPayload []byte

	resPayload, err := c.GetRequest(fmt.Sprintf("/render/%s/html", id))

	if err != nil {
		return nil, err
	}

	return &resPayload, nil
}
