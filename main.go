package main

import (
	"strings"
)

func main() {
	startRepl()
}

type cmds struct {
	name        string
	description string
}

var pokedex = map[string]cmds{
	"help": {
		name:        "help",
		description: "Displays a help message",
	},
	"exit": {
		name:        "exit",
		description: "Exits the pokedex",
	},
}

func cleanInput(input string) []string {
	lowered := strings.ToLower(input)
	args := strings.Fields(lowered)
	if len(args) < 1 {
		return nil
	}
	return args
}
