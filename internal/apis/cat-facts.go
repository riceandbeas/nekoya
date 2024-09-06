package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type CatFactApi struct {
	url    url.URL
	client *http.Client
}

func NewCatFactApi() (*CatFactApi, error) {
	return &CatFactApi{
		url: url.URL{
			Scheme: "https",
			Host:   "catfact.ninja",
		},
		client: http.DefaultClient,
	}, nil
}

func (api *CatFactApi) GetRandomFact() (string, error) {
	api.url.Path = "/fact"

	resp, err := api.client.Get(api.url.String())
	if err != nil {
		return "", fmt.Errorf("failed to get random fact: %w", err)
	}
	defer resp.Body.Close()

	var fact struct {
		Fact   string `json:"fact"`
		Length int    `json:"length"`
	}

	err = json.NewDecoder(resp.Body).Decode(&fact)
	if err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return fact.Fact, nil
}
