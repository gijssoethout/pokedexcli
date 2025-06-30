package main

import "fmt"

func commandExplore(cfg *config) error {
	if cfg.cmdArgs == nil {
		return fmt.Errorf("no given location found")
	}

	fmt.Printf("Exploring %s\n", cfg.cmdArgs[0])

	resp, err := cfg.pokeapiClient.GetPokemonPerLocation(cfg.cmdArgs[0])
	if err != nil {
		return err
	}

	for _, pokemon := range resp {
		fmt.Println(pokemon.Name)
	}

	return nil
}
