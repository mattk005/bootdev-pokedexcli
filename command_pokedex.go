package main

import (
	"fmt"
)

func commandPokedex(cfg *config, arg *string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("Your pokedex is empty")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, entry := range cfg.pokedex {
		fmt.Printf("- %s\n", entry.Name)
	}
	return nil
}
