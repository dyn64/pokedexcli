package main

import (
	"fmt"
	"math/rand"
)

func commandCatchPokemon(conf *config, args ...string) error {
	// makes sure the command has an argument
	if len(args) != 1 {
		return fmt.Errorf("Missing pokemon\n")
	}
	pokemon := args[0]

	// gets info about the pokemon
	pokemonInfo, err := conf.pokeapiClient.GetPokemonInfo(pokemon)
	if err != nil {
		return err
	}

	// tries to catch said pokemon
	// sets the difficulty to catch the pokemon based on their baseXP / 10
	difficulty := 100 - pokemonInfo.BaseExperience/10
	// rolls a d100
	attempt := rand.Intn(100)

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonInfo.Name)
	// checks if the pokemon is caught and adds it to the pokedex if it is caught
	if attempt <= difficulty {
		fmt.Printf("Caught %v!\n", pokemonInfo.Name)
		conf.pokeDex.Add(pokemonInfo)
	} else {
		fmt.Printf("%v escaped!\n", pokemonInfo.Name)
	}

	return nil
}
