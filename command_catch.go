package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, arg *string) error {
	resp, err := cfg.pokeapiClient.GetPokemon(arg)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", resp.Name)

	exp := resp.BaseExperience
	rand := rand.Intn(exp)
	if (700 - (rand + 100)) > 100 {
		fmt.Printf("%s was caught!\n", resp.Name)
		cfg.pokedex[resp.Name] = resp
	} else {
		fmt.Printf("%s escaped!\n", resp.Name)
	}
	// for k := range cfg.pokedex {
	// 	fmt.Println(cfg.pokedex[k].Name)
	// }
	return nil
}
