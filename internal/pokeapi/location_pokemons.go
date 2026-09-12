package pokeapi

import (
	"encoding/json"
)

// https://pokeapi.co/api/v2/location-area/{id or name}/

func (c *Client) ListLocationInfo(location string) (locationInfo, error) {
	// Contructs the URL used to fetch info from the PokeApi
	url := baseURL + "/location-area/" + location

	// Makes the httprequest
	data, err := c.MakeRequest(url, "GET")
	if err != nil {
		return locationInfo{}, err
	}

	// Unmarshals the data into a readable JSON-structure of type locationInfo{}
	locInfo := locationInfo{}
	err = json.Unmarshal(data, &locInfo)
	if err != nil {
		return locationInfo{}, nil
	}

	return locInfo, nil
}
