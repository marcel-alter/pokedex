package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

// Registry ↓↓↓↓↓↓
type config struct {
	commands map[string]cliCommand
	location Location
}

var configCom = config{
	commands: make(map[string]cliCommand),
	location: Location{PAGE: 0},
}

// Location structs ↓↓↓↓↓↓↓
type LocationArea struct {
	NAME string `json:"name"`
	URL  string `json:"url"`
}
type Location struct {
	PAGE     int            `json:"-"`
	COUNT    int            `json:"count"`
	NEXT     string         `json:"next"`
	PREVIOUS string         `json:"previous"`
	RESULTS  []LocationArea `json:"results"`
}

//var location = Location{PAGE: 0}

/*var Commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
}*/
