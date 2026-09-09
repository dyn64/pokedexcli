package main

import (
	"time"

	"github.com/dyn64/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	conf := &config{
		commands:      getCmd(),
		pokeapiClient: pokeClient,
	}
	startRepl(conf)
}
