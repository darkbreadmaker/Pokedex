package main

import (
	"fmt";
	"os"
)

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n \n")
	commandIndex := getCommands()
	for _, cmd := range commandIndex {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
