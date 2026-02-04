package main

import (
	"strings"
	"encoding/json"
	"bufio"
	"os"
	"fmt"
	"github.com/darkbreadmaker/Pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name string
	description string
	callback func( *config, ...string) error
}

var pokedex map[string]pokeapi.Pokemon = make(map[string]pokeapi.Pokemon)

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
		"explore": {
			name: "explore",
			description:"Explores a given area",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Throw a Pokeball",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Inspect a given Pokemon",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Lists all the Pokemon you've caught",
			callback: commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	var splitText []string = strings.Fields(lowerText)
	return splitText
}

func startRepl(cfg *config) {
	jsonData, err := os.ReadFile("myPokedex.json")
	if err != nil {
		fmt.Printf("Error parsing Pokedex data: %s\n", err)
	}
	err = json.Unmarshal(jsonData, &pokedex)
	if err != nil {
		fmt.Printf("Error reading Pokedex data: %s\n", err)
	}
	scanner := bufio.NewScanner(os.Stdin)
		for {
		fmt.Printf("Pokedex > ")
		 if scanner.Scan() == true {
			 input := scanner.Text()
			 cleaned := cleanInput(input)
		   firstWord := cleaned[0]
			 var param string
			 if len(cleaned) > 1 {
			 param = cleaned[1]
		 		}
			 commandIndex := getCommands()
			 cmd, ok := commandIndex[firstWord] 
			 if !ok {
				 fmt.Println("unknown command")
			 }else{
				 err := cmd.callback(cfg, param)
				 if err != nil {
					 fmt.Println(err)
				 }
			 }
		 }
	}
}
