package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config) error {
	if len(cfg.cmdArgs) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	name := cfg.cmdArgs[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	if rand.IntN(pokemon.BaseExperience) > 50 {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	fmt.Printf("%s was caught!\n", name)
	cfg.pokedex[name] = pokemon

	return nil
}
