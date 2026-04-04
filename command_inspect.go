package main

import "fmt"

func commandInspect(cfg *config, arg *string) error {
	entry, ok := cfg.pokedex[*arg]
	if !ok {
		println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %v\n", entry.Name)
	fmt.Printf("Height: %v\n", entry.Height)
	fmt.Printf("Weight: %v\n", entry.Weight)
	fmt.Println("Stats:")
	for _, stat := range entry.Stats {
		fmt.Printf("	-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, types := range entry.Types {
		fmt.Printf("	- %s\n", types.Type.Name)
	}
	return nil
}
