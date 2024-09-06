package apis

import (
	"encoding/json"
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

func (c *CatFactApi) GetRandomFact() (string, error) {
	c.url.Path = "/fact"

	resp, err := c.client.Get(c.url.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var fact struct {
		Fact   string `json:"fact"`
		Length int    `json:"length"`
	}

	err = json.NewDecoder(resp.Body).Decode(&fact)
	if err != nil {
		return "", err
	}

	return fact.Fact, nil
}
