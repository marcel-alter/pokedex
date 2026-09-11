package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.captured) == 0 {
		return fmt.Errorf("No Pokemon Caught yet!")
	}

	fmt.Print("Your Pokedex:\n")
	for _, pokemon := range cfg.captured {
		fmt.Println(pokemon.Name)
	}
	return nil
}
