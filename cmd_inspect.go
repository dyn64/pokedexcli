package main

import "fmt"

// Gets info about a specific pokemon from your pokedex
func commandInspectPokemon(conf *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("No pokemon specified")
	}
	pokemon := args[0]

	err := conf.pokeDex.Get(pokemon)
	if err != nil {
		return err
	}
	return nil
}
