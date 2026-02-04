package main

import (
	"fmt"
	"os"
	"encoding/json"
	"math/rand"
	"github.com/darkbreadmaker/Pokedex/internal/pokeapi"
)
func commandExit(cfg *config, param ...string) error {
	fmt.Printf("Saving your Pokedex\n")
	jsonData, err := json.Marshal(pokedex)
	if err != nil {
		fmt.Printf("Error saving Pokedex: %s\n", err)
	}
	err = os.WriteFile("myPokedex.json", jsonData, 0664)
	if err != nil {
		fmt.Printf("Error saving your Pokedex: %s\n", err)
	}
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, param ...string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n \n")
	commandIndex := getCommands()
	for _, cmd := range commandIndex {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
func commandMap(cfg *config, param ...string) error {
	var locationArea pokeapi.LocationArea
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

func commandMapb(cfg *config, param ...string) error {
	var locationArea pokeapi.LocationArea
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

func commandExplore(cfg *config, param ...string) error {
	if param[0] == "" {
		fmt.Printf("Please provide a location to explore\n")
	}else{
		url := "https://pokeapi.co/api/v2/location-area/" + param[0]
		locationData := cfg.pokeapiClient.ExploreLocation(url)
		pokemonList := locationData.PokemonEncounters
		fmt.Printf("Exploring %s...\n", param[0])
		fmt.Printf("Found Pokemon:\n") 
		for i := 0; i < len(pokemonList); i++ {
			fmt.Printf(" - %s\n", pokemonList[i].Pokemon.Name)
		}
	}
	return nil
}

func commandCatch(cfg *config, param ...string) error {	
	if param[0] == "" {
		fmt.Printf("Please provide a Pokemon to catch")
	}else{
		url := "https://pokeapi.co/api/v2/pokemon/" + param[0]
		fmt.Printf("Throwing a Pokeball at %s...\n", param[0])
		pokemon := cfg.pokeapiClient.CatchPokemon(url)
		pokemonName := pokemon.Name
		catchChance := pokemon.BaseExperience
		_, ok := pokedex[pokemonName]
		if ok {
				fmt.Printf("You already caught a %s\n", pokemonName)
			}else{
			if rand.Intn(650) > catchChance {
				fmt.Printf("%s was caught!\n", pokemonName)
				pokedex[pokemonName] = pokemon
			}else{
				fmt.Printf("%s escaped!\n", pokemonName)
			}
		}
	}
	return nil
}

func commandInspect(cfg *config, param ...string) error {
	if param[0] == "" {
		fmt.Println("Please provide a Pokemon to inspect")
	}else{
		pokemon, ok := pokedex[param[0]]
		if ok {
			fmt.Printf("Name: %s\n", pokemon.Name)
			fmt.Printf("Height: %d\n", pokemon.Height)
			fmt.Printf("Weight: %d\n", pokemon.Weight)
			fmt.Printf("Stats:\n")
			for i := 0; i < len(pokemon.Stats); i++ {
				statName := pokemon.Stats[i].Stat.Name
				statVal := pokemon.Stats[i].BaseStat + pokemon.Stats[i].Effort
				fmt.Printf("	-%s: %d\n", statName, statVal)
			}
			fmt.Printf("Types:\n")
			for i := 0; i < len(pokemon.Types); i++ {
				fmt.Printf("	- %s\n", pokemon.Types[i].Type.Name)
			}
		}else{
			fmt.Printf("You haven't caught %s yet\n", param[0])
		}
	}
	return nil
}

func commandPokedex(cfg *config, param ...string) error {
	if len(pokedex) == 0 {
		fmt.Printf("You haven't caught any Pokemon yet\n")
	}else{
		for key, _ := range pokedex {
			fmt.Printf(" - %s\n", key)
		}
	}
	return nil
}
