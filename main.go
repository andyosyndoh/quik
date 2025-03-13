package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: textindex [options]")
		os.Exit(1)
	}

	command := ""
	if os.Args[1] == "-c" {
		command = os.Args[2]
	} else {
		fmt.Println("Unknown command option")
		os.Exit(1)
	}
	args := os.Args[3:]
}
