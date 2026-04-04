package main

import (
	"time"

	"pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Second)
	pokedex := make(map[string]pokeapi.RespPokemonInfo)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       pokedex,
	}
	StartRepl(cfg)
}
