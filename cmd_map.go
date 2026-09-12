package main

import "fmt"

// Lists the (next 20) locations from the map and sets the next/prev urls in the config
func commandMapNext(conf *config, param ...string) error {
	locations, err := conf.pokeapiClient.ListLocations(conf.nextLocation)
	if err != nil {
		return err
	}

	conf.nextLocation = locations.Next
	conf.prevLocation = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

// Lists the previous 20 locations from the map. Prints an error if you are on the first page
// updates the next/prev urls in the config
func commandMapPrev(conf *config, param ...string) error {
	if conf.prevLocation == nil {
		return fmt.Errorf("your're on the first page")
	}

	locations, err := conf.pokeapiClient.ListLocations(conf.prevLocation)
	if err != nil {
		return err
	}

	conf.nextLocation = locations.Next
	conf.prevLocation = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
