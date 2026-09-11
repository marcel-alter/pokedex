package main

import (
	"time"

	"github.com/marcel-alter/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
		captured:      make(map[string]pokeapi.PokemonStruct),
	}

	startRepl(cfg)
}
