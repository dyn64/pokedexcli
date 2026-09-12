package main

import (
	"time"

	"github.com/dyn64/pokedexcli/internal/pokeapi"
)

// Initializes the (http) pokeClient and sets up the config with a client and a pokedex
func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	conf := &config{
		pokeapiClient: pokeClient,
		pokeDex:       pokeapi.NewPokedex(),
	}
	// starts the input loop
	startRepl(conf)
}
