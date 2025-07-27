package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	// TODO: add comment if 0 pokemon
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Println(" -", pokemon.Name)
	}

	return nil
}
