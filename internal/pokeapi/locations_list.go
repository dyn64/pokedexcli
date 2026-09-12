package pokeapi

import (
	"encoding/json"
)

func (c *Client) ListLocations(pageURL *string) (pokeMap, error) {
	// Contructs the URL used to fetch info from the PokeApi ..
	// .. unless its supplied to the function already
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Makes the httprequest
	data, err := c.MakeRequest(url, "GET")
	if err != nil {
		return pokeMap{}, err
	}

	// Unmarshals the data into a readable JSON-structure of type pokeMap{}
	locations := pokeMap{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return pokeMap{}, err
	}

	return locations, nil
}
