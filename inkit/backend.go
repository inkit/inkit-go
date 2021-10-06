package inkit

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"github.com/inkitio/gosdk/client"
)

func (c *client.Client) GetRequest(relUrl string) ([]byte, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.BaseUrl, relUrl), nil)

	if err != nil {
		return nil, err
	}

	res, err := c.sendRequest(req)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (c *client.Client) ListRequest(relUrl string, options *ListOptions) ([]byte, error) {
	pageSize := 50
	page := 1

	if options != nil {
		pageSize = options.PageSize
		page = options.Page
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s?page=%d&page_size=%d", c.BaseUrl, relUrl, page, pageSize), nil)

	if err != nil {
		return nil, err
	}

	res, err := c.sendRequest(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client.Client) PostRequest(relUrl string, payload []byte) ([]byte, error) {

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.BaseUrl, relUrl), bytes.NewBuffer(payload))

	if err != nil {
		return nil, err
	}

	/*if err := c.sendRequest(req, res); err != nil {
		return err
	}*/

	res, err := c.sendRequest(req)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (c *client.Client) sendRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Inkit-API-Token", c.apiKey)

	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		/*var errRes Response

		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}*/

		return nil, fmt.Errorf("unkown error, status code: %d", res.StatusCode)
		// handle error
	}
	
	var respBytes []byte

	respBytes, err = io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return respBytes, nil
}
type ListOptions struct {
	PageSize int
	Page int
}
