
package main

import (
	"fmt";
	"bufio";
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		 if scanner.Scan() == true {
			 input := scanner.Text()
			 cleaned := cleanInput(input)
		   firstWord := cleaned[0]
		 	 fmt.Printf("Your command was: %s \n", firstWord)
		 }
	}
}
