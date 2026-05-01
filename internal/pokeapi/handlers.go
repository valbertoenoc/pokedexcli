// Package pokeapi contains PokeAPI handlers
package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocation(pageURL *string) (ListLocationResponse, error) {
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
			return ListLocationResponse{}, fmt.Errorf("error fetching locations: %v", err)
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return ListLocationResponse{}, fmt.Errorf("error reading response data: %v", err)
		}

		c.cache.Add(url, body)
	}

	var locationResponse ListLocationResponse
	err := json.Unmarshal(body, &locationResponse)
	if err != nil {
		return ListLocationResponse{}, fmt.Errorf("error unserializing data: %v", err)
	}

	return locationResponse, nil
}

func (c *Client) GetLocationArea(area string) (Location, error) {
	url := BASE_URL + "location-area/" + area

	var body []byte
	if cacheEntry, ok := c.cache.Get(url); ok {
		body = cacheEntry
	} else {
		res, err := c.httpClient.Get(url)
		if err != nil {
			return Location{}, err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return Location{}, err
		}

		c.cache.Add(url, body)
	}

	var locationAreaPokemon Location
	err := json.Unmarshal(body, &locationAreaPokemon)
	if err != nil {
		return Location{}, err
	}

	return locationAreaPokemon, nil
}

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := BASE_URL + "pokemon/" + pokemonName

	var body []byte
	if cacheEntry, ok := c.cache.Get(url); ok {
		body = cacheEntry
	} else {
		res, err := c.httpClient.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}

		c.cache.Add(url, body)
	}

	var pokemon Pokemon
	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
