package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gijssoethout/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	pokedex          map[string]pokeapi.Pokemon
	cmdArgs          []string
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		cmdName := words[0]
		cmd, exists := getCommands()[cmdName]

		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		cfg.cmdArgs = []string{}
		if len(words) > 1 {
			cfg.cmdArgs = words[1:]
		}

		err := cmd.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
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
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"pokedex": {
			name:        "pokedex <pokemon_name>(optional)",
			description: "if a pokemon is given: displays information about a caught pokemon; otherwise displays list of caught pokemon",
			callback:    commandPokedex,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Displays all pokemon found in the given location",
			callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "Displays the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
