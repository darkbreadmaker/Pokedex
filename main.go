
package main

import (
	"github.com/darkbreadmaker/Pokedex/internal/pokeapi"
)
func main() {
	cfg := &config {
		pokeapiClient: pokeapi.NewClient(),
	}
	startRepl(cfg)
}


