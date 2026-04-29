// Package pokeapi contains PokeAPI handlers
package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageURL *string) (LocationListResponse, error) {
	url := BASE_URL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var body []byte
	cacheEntry, ok := c.cache.Get(url)
	if ok {
		body = cacheEntry
	} else {
		res, err := http.Get(url)
		if err != nil {
			return LocationListResponse{}, fmt.Errorf("error fetching locations: %v", err)
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationListResponse{}, fmt.Errorf("error reading response data: %v", err)
		}

		c.cache.Add(url, body)
	}

	var locationResponse LocationListResponse
	err := json.Unmarshal(body, &locationResponse)
	if err != nil {
		return LocationListResponse{}, fmt.Errorf("error unserializing data: %v", err)
	}

	return locationResponse, nil
}
