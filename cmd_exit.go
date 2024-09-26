package main

import (
	"fmt"
	"os"
)

func cmdExit() error {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}
