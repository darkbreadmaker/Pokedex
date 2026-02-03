package main

import (
	"strings";
	"bufio";
	"os";
	"fmt";
	"github.com/darkbreadmaker/Pokedex/internal/pokeapi"
)

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	var splitText []string = strings.Fields(lowerText)
	return splitText
}
type cliCommand struct {
	name string
	description string
	callback func( *config) error
}
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays next 20 locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays previous 20 locations",
			callback: commandMapb,
		},
	}
} 
type config struct {
	pokeapiClient pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}
func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
		for {
		fmt.Printf("Pokedex > ")
		 if scanner.Scan() == true {
			 input := scanner.Text()
			 cleaned := cleanInput(input)
		   firstWord := cleaned[0]
			 commandIndex := getCommands()
			 cmd, ok := commandIndex[firstWord] 
			 if !ok {
				 fmt.Println("unknown command")
			 }else{
				 err := cmd.callback(cfg)
				 if err != nil {
					 fmt.Println(err)
				 }
			 }
		 }
	}
}
