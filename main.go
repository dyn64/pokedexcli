package main

func main() {
	conf := &config{
		commands: getCmd(),
		mapNext:  "https://pokeapi.co/api/v2/location-area/",
		mapBack:  "https://pokeapi.co/api/v2/location-area/",
		mapBase:  "https://pokeapi.co/api/v2/location-area/",
	}
	startRepl(conf)
}
