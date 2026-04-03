package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(locationArea string) (LocationsDeep, error) {
	url := baseURL + "/location-area/" + locationArea

	body, exists := c.cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationsDeep{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationsDeep{}, err
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationsDeep{}, err
		}

		c.cache.Add(url, body)
	}

	locations := LocationsDeep{}
	err := json.Unmarshal(body, &locations)
	if err != nil {
		return LocationsDeep{}, err
	}

	return locations, nil

}
