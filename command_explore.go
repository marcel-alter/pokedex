package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	name := args[0]
	DeepLocationsResp, err := cfg.pokeapiClient.ListPokemon(name)
	if err != nil {
		//fmt.Print("wrong command")
		return err
	}

	for _, loc := range DeepLocationsResp.PokemonEncounters {
		fmt.Println("-", loc.Pokemon.Name)
	}
	return nil
}
