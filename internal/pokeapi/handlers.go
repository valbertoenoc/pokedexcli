// Package pokeapi contains PokeAPI handlers
package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetLocations(url string) (LocationListResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationListResponse{}, fmt.Errorf("error fetching locations: %v", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationListResponse{}, fmt.Errorf("error reading response data: %v", err)
	}

	var locationResponse LocationListResponse
	err = json.Unmarshal(data, &locationResponse)
	if err != nil {
		return LocationListResponse{}, fmt.Errorf("error unserializing data: %v", err)
	}

	return locationResponse, nil
}
