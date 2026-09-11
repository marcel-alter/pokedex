package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 && len(cfg.captured) == 0 {
		return fmt.Errorf("No Pokemon Caught yet to inspect!")
	} else if len(args) < 1 {
		fmt.Print("What Pokemon would you like to inspect? You have caught...\n")
		for _, pokemon := range cfg.captured {
			fmt.Println(pokemon.Name)
		}
		return nil
	}
	name := args[0]
	if pokemon, exists := cfg.captured[name]; exists == false {
		return fmt.Errorf("You have to catch %s first!", name)
	} else {
		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Height:", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)
		fmt.Printf(
			`Stats:
  -hp: %d
  -attack: %d
  -defense: %d
  -special-attack: %d
  -special-defense: %d
  -speed %d
`, pokemon.Stats[0].BaseStat, pokemon.Stats[1].BaseStat, pokemon.Stats[2].BaseStat,
			pokemon.Stats[3].BaseStat, pokemon.Stats[4].BaseStat, pokemon.Stats[5].BaseStat,
		)
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		}

		return nil
	}
}
