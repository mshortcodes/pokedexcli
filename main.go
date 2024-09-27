package main

import (
	"pokedexcli/internal/pokeapi"
	"time"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextURL       *string
	prevURL       *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}

	startRepl(&cfg)
}
