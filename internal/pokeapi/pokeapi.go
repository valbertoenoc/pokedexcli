// Package pokeapi contains PokeAPI definitions
package pokeapi

const (
	BASE_URL = "https://pokeapi.co/api/v2/"
)

type LocationListResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
