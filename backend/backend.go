package backend

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	BaseURLV1 = "https://api.inkit.com/v1"
)

type Backend struct {
	BaseUrl    string
	apiKey     string
	httpClient *http.Client
}

func NewBackend(apiKey string) *Backend {
	return &Backend{
		BaseUrl: BaseURLV1,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (b *Backend) GetRequest(relUrl string) ([]byte, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", b.BaseUrl, relUrl), nil)

	if err != nil {
		return nil, err
	}

	res, err := b.sendRequest(req)

	if err != nil {
		return nil, err
	}

	return res, nil

}

type ListOptions struct {
	Page     int
	PageSize int
}

func (b *Backend) ListRequest(relUrl string, options *ListOptions) ([]byte, error) {
	pageSize := 50
	page := 1

	if options != nil {
		pageSize = options.PageSize
		page = options.Page
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s?page=%d&page_size=%d", b.BaseUrl, relUrl, page, pageSize), nil)

	if err != nil {
		return nil, err
	}

	res, err := b.sendRequest(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (b *Backend) PostRequest(relUrl string, payload []byte) ([]byte, error) {

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", b.BaseUrl, relUrl), bytes.NewBuffer(payload))

	if err != nil {
		return nil, err
	}

	/*if err := c.sendRequest(req, res); err != nil {
		return err
	}*/

	res, err := b.sendRequest(req)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (b *Backend) sendRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Inkit-API-Token", b.apiKey)

	res, err := b.httpClient.Do(req)

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
	Page     int
}
