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

func NewHttpCatsApi() *HttpCatsApi {
	return &HttpCatsApi{
		url: url.URL{
			Scheme: "https",
			Host:   "http.cat",
		},
		client: http.DefaultClient,
	}
}

func (api *HttpCatsApi) GetStatusImage(status string) (string, error) {
	api.url.Path = status

	resp, err := api.client.Get(api.url.String())
	if err != nil {
		return "", fmt.Errorf("failed to get status image: %w", err)
	}
	defer resp.Body.Close()

	// user probably entered an invalid status code
	if resp.StatusCode != http.StatusOK {
		api.url.Path = "404"
		return api.url.String(), nil
	}

	return api.url.String(), nil
}
