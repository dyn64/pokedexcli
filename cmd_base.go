package main

import (
	"fmt"
	"os"
)

func commandExit(conf *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config, args ...string) error {
	commands := getCmd()
	if len(args) != 0 {
		cmd, ok := commands[args[0]]
		if ok {
			fmt.Printf("%s: %s\n", cmd.name, cmd.description)
			return nil
		}
	}
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage: \"help <command>\" (optional) \n\n")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
