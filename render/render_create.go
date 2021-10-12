package render

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (c *RenderClient) CreateRender(options *CreateRenderOptions) (*Render, error) {
	render := &Render{}
	var resPayload []byte

	if options == nil {
		return nil, fmt.Errorf("must provide CreateRenderOptions to create a render")
	}

	if options.Html == "" && options.FileName == "" {
		return nil, fmt.Errorf("must provide html or a file to create a render")
	}

	if options.Html != "" && options.FileName != "" {
		return nil, fmt.Errorf("either Html or FileName can be provided")
	}

	if options.FileName != "" {

		file, err := ioutil.ReadFile(options.FileName)

		if err != nil {
			return nil, err
		}
		options.Html = string(file)
	}

	options.Html = base64.StdEncoding.EncodeToString([]byte(options.Html))

	payload, err := json.Marshal(options)

	if err != nil {
		return nil, err
	}

	resPayload, err = c.Backend.PostRequest("/render", payload)

	if err != nil {
		return nil, err
	}

	if err = json.NewDecoder(bytes.NewBuffer(resPayload)).Decode(render); err == nil {
		return render, nil
	} else {
		return nil, err
	}
}
