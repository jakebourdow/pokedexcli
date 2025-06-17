package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	availableCommands := getCommands()
	if len(availableCommands) == 0 {
		return fmt.Errorf("no commands available")
	}
	for i := range availableCommands {
		fmt.Printf("%s: %s\n", availableCommands[i].name, availableCommands[i].description)
	}
	return nil
}
