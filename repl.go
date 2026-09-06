package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// var commands = map[string]cliCommand{}

func getCmd() map[string]cliCommand {
	// commands = make(map[string]cliCommand)
	// commands["exit"] = cliCommand{
	// 	name:        "exit",
	// 	description: "Exit the Pokedex",
	// 	callback:    commandExit,
	// }
	// commands["help"] = cliCommand{
	// 	name:        "help",
	// 	description: "Displays a help message",
	// 	callback:    commandHelp,
	// }
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
	}
}

func startRepl() {

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
			cmd.callback()
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
