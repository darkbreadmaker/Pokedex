
package main

import (
	"fmt";
	"bufio";
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func() error
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
	}
} 
func main() {
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
				 err := cmd.callback()
				 if err != nil {
					 fmt.Println(err)
				 }
			 }
		 }
	}
}


