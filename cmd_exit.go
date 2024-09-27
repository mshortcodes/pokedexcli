package main

import (
	"fmt"
	"os"
)

func cmdExit(cfg *config, args ...string) error {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}
