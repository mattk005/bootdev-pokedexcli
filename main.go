package main

import (
	"time"

	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cache := pokecache.NewCache(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokecache:     cache,
	}
	StartRepl(cfg)
}
