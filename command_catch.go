package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("what Pokemon do you want to catch?")
	}
	name := args[0]
	if cfg.captured == nil {
		return fmt.Errorf("config.captured still empty!!!")
	}
	if _, exists := cfg.captured[name]; exists {
		return fmt.Errorf("%s already caught!?!\n", name)
	}
	PokemonResp, err := cfg.pokeapiClient.PokemonFind(name)
	if err != nil {
		//fmt.Print("wrong command")
		return err
	}
	//fmt.Printf("%d chance of success! %d baseEX\n", rand.Intn(PokemonResp.BaseExperience), PokemonResp.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	chance := rand.Intn(PokemonResp.BaseExperience)
	//fmt.Print(chance)
	if chance < 50 {
		fmt.Printf("%s was caught!\n", name)
		cfg.captured[name] = PokemonResp
		return nil
	} else {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}
}
