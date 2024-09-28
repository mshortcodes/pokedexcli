package main

import "fmt"

func cmdPokedex(cfg *config, args ...string) error {
	fmt.Println("Your pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
