package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type TheCatApi struct {
	url    url.URL
	client *http.Client
}

func NewTheCatApi() (*TheCatApi, error) {
	return &TheCatApi{
		url: url.URL{
			Scheme: "https",
			Host:   "api.thecatapi.com",
		},
		client: http.DefaultClient,
	}, nil
}

func (api *TheCatApi) GetRandomImage(breed string) (string, error) {
	if breed != "" {
		breedId, err := api.getBreedId(strings.ToLower(breed))
		if err != nil {
			return "", fmt.Errorf("failed to get breed id: %w", err)
		}

		api.url.RawQuery = fmt.Sprintf("breed_ids=%s", breedId)
	}

	api.url.Path = "/v1/images/search"

	resp, err := api.client.Get(api.url.String())
	if err != nil {
		return "", fmt.Errorf("failed to get random image: %w", err)
	}
	defer resp.Body.Close()

	var image []struct {
		Url string `json:"url"`
	}

	err = json.NewDecoder(resp.Body).Decode(&image)
	if err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return image[0].Url, nil
}

func (api *TheCatApi) GetBreeds() (map[string]string, error) {
	api.url.Path = "/v1/breeds"

	resp, err := api.client.Get(api.url.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get breeds: %w", err)
	}
	defer resp.Body.Close()

	var breeds []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	err = json.NewDecoder(resp.Body).Decode(&breeds)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	breedsMap := make(map[string]string)
	for _, breed := range breeds {
		breedsMap[strings.ToLower(breed.Name)] = breed.Id
	}

	return breedsMap, nil
}

func (api *TheCatApi) getBreedId(breed string) (string, error) {
	breeds, err := api.GetBreeds()
	if err != nil {
		return "", fmt.Errorf("failed to get breeds: %w", err)
	}

	breedId, ok := breeds[breed]
	if !ok {
		return "", fmt.Errorf("unknown breed: %s", breed)
	}

	return breedId, nil
}
