package pokeapi

import (
	"encoding/json"
)

// https://pokeapi.co/api/v2/pokemon/{id or name}/

func (c *Client) GetPokemonInfo(pokemon string) (pokemonInfo, error) {
	url := baseURL + "/pokemon/" + pokemon

	// Makes the httprequest
	data, err := c.MakeRequest(url, "GET")
	if err != nil {
		return pokemonInfo{}, err
	}

	// Unmarshals the data into a readable JSON-structure of type pokemonInfo{}
	pokInfo := pokemonInfo{}
	err = json.Unmarshal(data, &pokInfo)
	if err != nil {
		return pokemonInfo{}, nil
	}
	return pokInfo, nil
}
