package pokeapi

import "fmt"

// Creates a new Pokedex
func NewPokedex() PokeDex {
	return PokeDex{
		dex: make(map[string]pokemonInfo),
	}
}

// PokeDex struct containing just a map with data of type pokemoninfo
type PokeDex struct {
	dex map[string]pokemonInfo
}

// Adds an entry into the pokedex
func (d *PokeDex) Add(pokemon pokemonInfo) error {
	// checks if you already have the pokemon
	_, ok := d.dex[pokemon.Name]
	if ok {
		fmt.Printf("You already have a %s\n", pokemon.Name)
		return nil
	}
	d.dex[pokemon.Name] = pokemon
	return nil
}

// Gets info about a pokemon in the pokedex
func (d *PokeDex) Get(pokemon string) error {

	// Checks if you have data about the pokemon requested (caught it before)
	pokeinfo, ok := d.dex[pokemon]
	if !ok {
		return fmt.Errorf("you have not caught %v yet", pokemon)
	}
	// Just prints out the info
	fmt.Printf("Name: %v\n", pokeinfo.Name)
	fmt.Printf("Height: %v\n", pokeinfo.Height)
	fmt.Printf("Weight: %v\n", pokeinfo.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokeinfo.Stats {
		fmt.Printf("\t-%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, Type := range pokeinfo.Types {
		fmt.Printf("\t- %v\n", Type.Type.Name)
	}
	fmt.Println()
	return nil
}

// Lists all the pokemon in your pokedex
func (d *PokeDex) List() error {
	// Makes sure your pokedex has any entries
	if len(d.dex) == 0 {
		return fmt.Errorf("You have not caught any pokemon!")
	}
	fmt.Println("Your Pokedex:")
	for _, poke := range d.dex {
		fmt.Printf("\t- %v\n", poke.Name)
	}
	return nil
}
