package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var cmd = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandhelp,
	},
	"exit": {
		name:        "exit",
		description: "",
		callback:    commandExit,
	},
}

func commandhelp() error {
	//Todo
	return nil
}

func commandExit() error {
	//Todo
	return nil
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex! \nUsage:")
	for _, command := range cmd {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	fmt.Print("pokedex >")

	for scanner.Scan() {
		fmt.Print("pokedex >")
		// input, ok := cmd[scanner.Text()]
		// if ok {
		// 	fmt.Printf("")
		// }
	}
}
