package main

import "fmt"

func commandExplore(cfg *config, argument *string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocationArea(argument)
	if err != nil {
		return err
	}
	for _, pokemon := range locationsResp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
