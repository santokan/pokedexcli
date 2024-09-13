package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, info := range getCommands() {
		fmt.Printf("  %s: %s\n", info.name, info.description)
	}
	fmt.Println()
	return nil
}
