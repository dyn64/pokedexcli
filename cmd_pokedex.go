package main

func commandPokedex(conf *config, args ...string) error {
	err := conf.pokeDex.List()
	if err != nil {
		return err
	}
	return nil
}
