package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"pokedexcli/internal/pokecache"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string, cache *pokecache.Cache) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	// Add cache block
	if dat, ok := cache.Get(url); ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(dat, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		// fmt.Println("---Used Cache---")
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	// if it's not cached, add it to cache
	cache.Add(url, dat)
	if err != nil {
		return RespShallowLocations{}, err
	}
	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}
	return locationsResp, nil
}
