package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// https://pokeapi.co/api/v2/location-area/{id or name}/

func (c *Client) ListLocationInfo(location string) (LocationInfo, error) {
	url := baseURL + "/location-area/" + location

	data, ok := c.pokeCache.Get(url)

	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationInfo{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationInfo{}, err
		}

		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationInfo{}, nil
		}
		c.pokeCache.Add(url, data)
	}
	locInfo := LocationInfo{}
	err := json.Unmarshal(data, &locInfo)
	if err != nil {
		return LocationInfo{}, nil
	}

	return locInfo, nil
}
