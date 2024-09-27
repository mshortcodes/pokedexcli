package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	pokedex := getCmds()

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}
		command := cleaned[0]
		cmd, ok := pokedex[command]
		if !ok {
			fmt.Println("Not a valid command.")
			fmt.Println()
			continue
		}
		err := cmd.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(input string) []string {
	lowered := strings.ToLower(input)
	args := strings.Fields(lowered)
	if len(args) < 1 {
		return nil
	}
	return args
}
