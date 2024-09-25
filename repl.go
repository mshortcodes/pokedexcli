package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}
		cmd := cleaned[0]
		switch cmd {
		case "help":
			fmt.Println(pokedex[cmd].description)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Not a valid command.")
		}
	}
}
