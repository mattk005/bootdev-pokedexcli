package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetPokemon
func (c *Client) GetPokemon(argument *string) (RespPokemonInfo, error) {
	if argument == nil {
		fmt.Println("No argument (pokemon) for client.getPokemon()")
		return RespPokemonInfo{}, nil
	}
	url := baseURL + fmt.Sprintf("/pokemon/%s/", *argument)
	if dat, ok := c.cache.Get(url); ok {
		pokemonResp := RespPokemonInfo{}
		err := json.Unmarshal(dat, &pokemonResp)
		if err != nil {
			return pokemonResp, err
		}
		return pokemonResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonInfo{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonInfo{}, err
	}
	if resp.StatusCode > 299 {
		return RespPokemonInfo{}, fmt.Errorf("Invalid pokemon, error code :%d\n", resp.StatusCode)
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemonInfo{}, err
	}
	c.cache.Add(url, dat)
	pokemonResp := RespPokemonInfo{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespPokemonInfo{}, err
	}
	return pokemonResp, nil
}
