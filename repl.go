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
	callback    func(*config) error
}

type config struct {
	commands      map[string]cliCommand
	pokeapiClient pokeapi.Client
	nextLocation  *string
	prevLocation  *string
}

func getCmd() map[string]cliCommand {

	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
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
		cmd, exists := getCmd()[input[0]]
		if exists {
			cmd.callback(conf)
		} else {
			fmt.Printf("Unknown command: %s\n", input[0])
		}

	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	result := strings.Fields(text)
	return result
}
