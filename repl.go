package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokecache"
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
		commandName := userInput[0]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
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
	pokecache        *pokecache.Cache
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
	}
}
