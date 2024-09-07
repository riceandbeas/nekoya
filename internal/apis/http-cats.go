package apis

import (
	"fmt"
	"net/http"
	"net/url"
)

type HttpCatsApi struct {
	url    url.URL
	client *http.Client
}

func NewHttpCatsApi() (*HttpCatsApi, error) {
	return &HttpCatsApi{
		url: url.URL{
			Scheme: "https",
			Host:   "http.cat",
		},
		client: http.DefaultClient,
	}, nil
}

func (api *HttpCatsApi) GetStatusImage(status string) (string, error) {
	api.url.Path = status

	resp, err := api.client.Get(api.url.String())
	if err != nil {
		return "", fmt.Errorf("failed to get status image: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get status image: %s", resp.Status)
	}

	return api.url.String(), nil
}
