package main

import (
	"fmt"
	"os"
)

func cmdExit(cfg *config) error {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}
