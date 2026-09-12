package main

// Lists all the pokemon in your pokedex
func commandPokedex(conf *config, args ...string) error {
	err := conf.pokeDex.List()
	if err != nil {
		return err
	}
	return nil
}
