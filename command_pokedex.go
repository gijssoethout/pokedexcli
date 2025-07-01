package main

import "fmt"

func commandPokedex(cfg *config) error {
	if len(cfg.cmdArgs) == 0 {
		fmt.Println("Your Pokedex:")

		for key := range cfg.pokedex {
			fmt.Printf("  - %s\n", key)
		}

		fmt.Println()
		fmt.Println("For more information about a pokemon use: pokedex <name_pokemon>")

		return nil
	}

	pokemon, ok := cfg.pokedex[cfg.cmdArgs[0]]
	if !ok {
		fmt.Printf("you have not caught %s\n", cfg.cmdArgs[0])
		return nil
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}

	return nil
}
