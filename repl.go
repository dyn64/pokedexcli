package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dyn64/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	commands      map[string]cliCommand
	pokeapiClient pokeapi.Client
	nextLocation  *string
	prevLocation  *string
	pokeDex       pokeapi.PokeDex
}

func getCmd() map[string]cliCommand {

	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help <command>",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists map locations",
			callback:    commandMapNext,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous 20 map locations",
			callback:    commandMapPrev,
		},
		"explore": {
			name:        "explore <location/id>",
			description: "Lists pokemon in specified location or id",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Tries to catch the pokemon specified",
			callback:    commandCatchPokemon,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Inspect specified Pokemon if you have caught it",
			callback:    commandInspectPokemon,
		},
	}
}

func startRepl(conf *config) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}

		cmd, exists := getCmd()[commandName]
		if exists {
			err := cmd.callback(conf, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", commandName)
		}

	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	result := strings.Fields(text)
	return result
}
