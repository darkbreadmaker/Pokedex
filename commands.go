package main

import (
	"fmt";
	"os";
	"github.com/darkbreadmaker/Pokedex/internal/pokeapi"
)
func commandExit(cfg *config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n \n")
	commandIndex := getCommands()
	for _, cmd := range commandIndex {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
func commandMap(cfg *config) error {
	var locationArea pokeapi.Location
	if cfg.nextLocationsURL == nil {
		locationArea = cfg.pokeapiClient.GetLocationArea("https://pokeapi.co/api/v2/location-area/")
	}else {
		locationArea = cfg.pokeapiClient.GetLocationArea(*cfg.nextLocationsURL)
	}
	for i := 0; i < len(locationArea.Results); i++ {
		fmt.Printf("%s\n",locationArea.Results[i].Name)
	}
	cfg.nextLocationsURL = locationArea.Next
	cfg.prevLocationsURL = locationArea.Previous
	return nil
}

func commandMapb(cfg *config) error {
	var locationArea pokeapi.Location
	if cfg.prevLocationsURL == nil {
		fmt.Printf("you're on the first page")
		return nil
	}else {
		locationArea = cfg.pokeapiClient.GetLocationArea(*cfg.prevLocationsURL)
	}
	for i := 0; i < len(locationArea.Results); i++ {
		fmt.Printf("%s\n",locationArea.Results[i].Name)
	}
	cfg.nextLocationsURL = locationArea.Next
	cfg.prevLocationsURL = locationArea.Previous
	return nil
}

