package main

func main() {
	conf := &config{
		commands: getCmd(),
	}
	startRepl(conf)
}
