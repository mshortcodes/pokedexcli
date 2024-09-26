package main

import (
	"fmt"
	"log"
)

func cmdMap(cfg *config) error {
	areas, err := cfg.pokeapiClient.ListAreas()
	if err != nil {
		log.Fatal(err)
	}
	for _, area := range areas.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextURL = areas.Next
	cfg.prevURL = areas.Previous
	return nil
}
