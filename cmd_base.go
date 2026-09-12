package main

import (
	"fmt"
	"os"
)

// command for exiting the program
func commandExit(conf *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// Lists help for command(s)
func commandHelp(conf *config, args ...string) error {
	// Gets the list of commands
	commands := getCmd()

	// Checks for an argument. For example "help map"
	// prints only the help for that command and returns
	if len(args) != 0 {
		cmd, ok := commands[args[0]]
		if ok {
			fmt.Printf("%s: %s\n", cmd.name, cmd.description)
			return nil
		}
	}
	// prints all the commands with info
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage: \"help <command>\" (optional) \n\n")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
