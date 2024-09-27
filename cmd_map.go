package main

import (
	"errors"
	"fmt"
)

func cmdMap(cfg *config, args ...string) error {
	areas, err := cfg.pokeapiClient.ListAreas(cfg.nextURL)
	if err != nil {
		return err
	}
	for _, area := range areas.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextURL = areas.Next
	cfg.prevURL = areas.Previous
	return nil
}

func cmdMapb(cfg *config, args ...string) error {
	if cfg.prevURL == nil {
		return errors.New("you are on the first page")
	}
	areas, err := cfg.pokeapiClient.ListAreas(cfg.prevURL)
	if err != nil {
		return err
	}
	for _, area := range areas.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextURL = areas.Next
	cfg.prevURL = areas.Previous
	return nil
}
