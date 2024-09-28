package main

import (
	"pokedexcli/internal/pokeapi"
	"time"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextURL       *string
	prevURL       *string
	caughtPokemon map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}
