package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"pokedexcli/internal/pokeapi"
)

func StartRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		userInput := cleanInput(reader.Text())
		if len(userInput) == 0 {
			continue
		}
		var argument string
		commandName := userInput[0]
		if len(userInput) > 1 {
			argument = userInput[1]
		}
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, &argument)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	fields := strings.Fields(lowerCase)
	return fields
}

type config struct {
	pokeapiClient    pokeapi.Client
	pokedex          map[string]pokeapi.RespPokemonInfo
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "prints a help message describing how to use the REPL",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exits the program",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations, and so on. This will be how we explore the Pokemon world",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the previous 20 locations. If on page 0, it will display an error message",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Takes an area as an argument and returns a list of all pokemon in that area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a Pokemon (\"throwing Pokeball at ...\") chance is based on pokemon's base exp, caught pokemon is added to pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "takes the name of a Pokemon and prints the name, height, weight, stats and type(s)",
			callback:    commandInspect,
		},
	}
}
