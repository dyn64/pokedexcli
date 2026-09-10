package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (pokeMap, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, ok := c.pokeCache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return pokeMap{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return pokeMap{}, err
		}

		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return pokeMap{}, err
		}
		c.pokeCache.Add(url, data)
	}
	locations := pokeMap{}
	err := json.Unmarshal(data, &locations)
	if err != nil {
		return pokeMap{}, err
	}

	return locations, nil
}
