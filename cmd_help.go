package main

import "fmt"

func cmdHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	pokedex := getCmds()
	for _, v := range pokedex {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	fmt.Println()
	return nil
}
