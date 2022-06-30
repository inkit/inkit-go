package render

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (c *RenderClient) Create(options *CreateRenderOptions) (*Render, error) {
	render := &Render{}
	var resPayload []byte

	if options == nil {
		return nil, fmt.Errorf("must provide CreateRenderOptions to create a render")
	}
	// No XOR in golang
	
	count := 0
	
	if options.Html != "" {
		count++
	}
	
	if options.FileName != "" {
		count++
	}
	
	if options.TemplateId != "" {
		count++
	}
	
	if count == 0 {
		return nil, fmt.Errorf("Must provide Html, FileName, or TemplateId to create a render")
	}

	if count > 1 {
		return nil, fmt.Errorf("Only one of the following can be provided at a time: Html, FileName, or TemplateId")
	}
	
	if options.TemplateId == "" {
		if options.FileName != "" {

			file, err := ioutil.ReadFile(options.FileName)

			if err != nil {
				return nil, err
			}
			options.Html = string(file)
		}

		options.Html = base64.StdEncoding.EncodeToString([]byte(options.Html))
	}

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
