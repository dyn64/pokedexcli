package pokeapi

import "fmt"

type PokeDex struct {
	dex map[string]Pokemon
}

func NewPokedex() PokeDex {
	dex := make(map[string]Pokemon)
	return PokeDex{
		dex: dex,
	}
}

func (d *PokeDex) Add(pokemon pokemonInfo) error {
	_, ok := d.dex[pokemon.Name]
	if ok {
		fmt.Printf("You already have a %s\n", pokemon.Name)
		return nil
	}
	d.dex[pokemon.Name] = Pokemon{
		Name:     pokemon.Name,
		Level:    1,
		pokeInfo: pokemon,
	}
	return nil
}
