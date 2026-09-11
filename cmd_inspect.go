package main

import "fmt"

func commandInspectPokemon(conf *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("No pokemon specified")
	}
	pokemon := args[0]
	//pokeinfo, ok := conf.pokeDex[]
	// pokeinfo, ok := conf.pokeDex.dex[pokemon]
	err := conf.pokeDex.Get(pokemon)
	if err != nil {
		return err
	}
	return nil
}
