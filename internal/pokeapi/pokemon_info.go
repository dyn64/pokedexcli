package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// https://pokeapi.co/api/v2/pokemon/{id or name}/

func (c *Client) GetPokemonInfo(pokemon string) (pokemonInfo, error) {
	url := baseURL + "/pokemon/" + pokemon

	data, ok := c.pokeCache.Get(url)

	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return pokemonInfo{}, nil
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return pokemonInfo{}, nil
		}

		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return pokemonInfo{}, nil
		}
		c.pokeCache.Add(url, data)
	}

	pokInfo := pokemonInfo{}
	err := json.Unmarshal(data, &pokInfo)
	if err != nil {
		return pokemonInfo{}, nil
	}
	return pokInfo, nil
}
