package pokeapi

import "fmt"

func NewPokedex() PokeDex {
	return PokeDex{
		dex: make(map[string]pokemonInfo),
	}
}

type PokeDex struct {
	dex map[string]pokemonInfo
}

func (d *PokeDex) Add(pokemon pokemonInfo) error {
	_, ok := d.dex[pokemon.Name]
	if ok {
		fmt.Printf("You already have a %s\n", pokemon.Name)
		return nil
	}
	// d.dex[pokemon.Name] = Pokemon{
	// 	Name:     pokemon.Name,
	// 	Level:    1,
	// 	pokeInfo: pokemon,
	// }
	d.dex[pokemon.Name] = pokemon
	return nil
}

func (d *PokeDex) Get(pokemon string) error {
	pokeinfo, ok := d.dex[pokemon]
	if !ok {
		return fmt.Errorf("you have not caught %v yet", pokemon)
	}
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

func (d *PokeDex) List() error {
	if len(d.dex) == 0 {
		return fmt.Errorf("You have not caught any pokemon!")
	}
	fmt.Println("Your Pokedex:")
	for _, poke := range d.dex {
		fmt.Printf("\t- %v\n", poke.Name)
	}
	return nil
}
