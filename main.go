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
		description: "Exits the application",
		callback:    commandExit,
	},
}

var userInput = make(chan string)

func getInput() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		userInput <- scanner.Text()
	}
}

func commandhelp() error {
	//TODO
	// should send the list of commands back
	return nil
}

func commandExit() error {
	//TODO
	// what this command should do when comes in as an input
	return nil
}

func main() {

	go getInput()

	fmt.Println("Welcome to the Pokedex! \nUsage:")
	for _, command := range cmd {
		fmt.Printf("\n%s: %s\n", command.name, command.description)
	}
	fmt.Print("pokedex > ")

	for input := range userInput {
		fmt.Print("pokedex > ")
		switch input {
		case "help":
			for _, command := range cmd {
				fmt.Printf("\n%s: %s\n", command.name, command.description)
			}
		case "exit":
			os.Exit(1)
		}
		fmt.Print("pokedex > ")
	}

}
