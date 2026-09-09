package main

import (
	"time"

	"github.com/marcel-alter/pokedexcli/internal/pokeapi"
	//"github.com/marcel-alter/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
