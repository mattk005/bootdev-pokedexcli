package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListLocationArea -
func (c *Client) ListLocationArea(argument *string) (RespDeepLocations, error) {
	if argument == nil {
		fmt.Println("No URL for client.ListLocationArea()")
		return RespDeepLocations{}, nil
	}
	url := baseURL + fmt.Sprintf("/location-area/%s/", *argument)
	// Add cache block
	if dat, ok := c.cache.Get(url); ok {
		locationsResp := RespDeepLocations{}
		err := json.Unmarshal(dat, &locationsResp)
		if err != nil {
			return RespDeepLocations{}, err
		}
		// fmt.Println("---Used Cache---")
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDeepLocations{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDeepLocations{}, err
	}
	if resp.StatusCode > 299 {
		fmt.Printf("Invalid area, error code :%d\n", resp.StatusCode)
		return RespDeepLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDeepLocations{}, err
	}
	// if it's not cached, add it to cache
	c.cache.Add(url, dat)
	locationsResp := RespDeepLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespDeepLocations{}, err
	}
	return locationsResp, nil
}
