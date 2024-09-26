package main

type cmd struct {
	name        string
	description string
	callback    func(*config) error
}

func getCmds() map[string]cmd {
	return map[string]cmd{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    cmdHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			callback:    cmdExit,
		},
		"map": {
			name:        "map",
			description: "lists location areas",
			callback:    cmdMap,
		},
	}
}
