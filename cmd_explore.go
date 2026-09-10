package main

import "fmt"

func commandExplore(conf *config, args ...string) error {
	if len(args) != 1 {
		fmt.Printf("Missing location\n")
		return nil
	}
	location := args[0]

	locationInfo, err := conf.pokeapiClient.ListLocationInfo(location)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationInfo.Name)
	fmt.Println("Found Pokemon: ")
	for _, pokemans := range locationInfo.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemans.Pokemon.Name)
	}
	return nil
}
