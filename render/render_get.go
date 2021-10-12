package render

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func (c *RenderClient) Get(id string) (*Render, error) {
	render := &Render{}

	var resPayload []byte
	resPayload, err := c.Backend.GetRequest(fmt.Sprintf("/render/%s", id))

	if err != nil {
		return nil, err
	}

	if err = json.NewDecoder(bytes.NewBuffer(resPayload)).Decode(render); err == nil {
		return render, nil
	} else {
		return nil, err
	}
}
func (c *RenderClient) GetPdfAndSaveToFile(id string, fileName string) error {
	pdfData, err := c.GetPdf(id)

	if err != nil {
		return err
	}

	f, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer f.Close()
	_, err = f.Write(*pdfData)

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (c *RenderClient) GetPdf(id string) (*[]byte, error) {
	var resPayload []byte

	resPayload, err := c.Backend.GetRequest(fmt.Sprintf("/render/%s/pdf", id))

	if err != nil {
		return nil, err
	}

	return &resPayload, nil
}

func (c *RenderClient) GetHtml(id string) (*[]byte, error) {
	var resPayload []byte

	resPayload, err := c.Backend.GetRequest(fmt.Sprintf("/render/%s/html", id))

	if err != nil {
		return nil, err
	}

	return &resPayload, nil
}

func (c *RenderClient) GetHtmlAndSaveToFile(id string, fileName string) error {
	htmlData, err := c.GetHtml(id)

	if err != nil {
		return err
	}

	f, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer f.Close()
	_, err = f.Write(*htmlData)

	if err != nil {
		return err
	} else {
		return nil
	}
}
