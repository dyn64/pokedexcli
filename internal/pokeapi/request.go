package pokeapi

import (
	"io"
	"net/http"
)

func (c *Client) MakeRequest(url string, request_type string) ([]byte, error) {
	// Polls the pokecache to see if we have requested the same url lately
	data, ok := c.pokeCache.Get(url)
	// If not, sends a new request to the PokeApi
	if !ok {
		// "Constructs" the httprequest, and returns the error message if it fails
		req, err := http.NewRequest(request_type, url, nil)
		if err != nil {
			return nil, err
		}
		// sends the httprequest to the client. Returns the error on failure
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		// makes sure the function closes the response before returning
		defer resp.Body.Close()
		// reads the body from the response into something half-useful
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		// adds the response to the cache
		c.pokeCache.Add(url, data)
	}
	return data, nil
}
