package main

import (
	"fmt"
	"math/rand"
)

func commandCatchPokemon(conf *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Missing pokemon")
	}
	pokemon := args[0]

	pokemonInfo, err := conf.pokeapiClient.GetPokemonInfo(pokemon)
	if err != nil {
		return err
	}

	difficulty := 100 - pokemonInfo.BaseExperience/10
	attempt := rand.Intn(100)
	// For debugging
	//	fmt.Printf("Chance to catch \"%s\": %d\nAttempt: %v\n", pokemonInfo.Name, difficulty, attempt)
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonInfo.Name)
	if attempt <= difficulty {
		fmt.Printf("Caught %v!\n", pokemonInfo.Name)
		conf.pokeDex.Add(pokemonInfo)
	} else {
		fmt.Printf("%v escaped!\n", pokemonInfo.Name)
	}

	return nil
}
