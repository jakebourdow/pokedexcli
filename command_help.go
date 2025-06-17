package main

import "fmt"

func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	availableCommands := getCommands()
	if len(availableCommands) == 0 {
		return fmt.Errorf("no commands available")
	}
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
