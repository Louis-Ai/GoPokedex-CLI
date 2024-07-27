package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
		"map": {
			name:        "map",
			description: "The map command displays the names of 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Similar to the map command, however, it displays the previous 20 locations.",
			callback:    commandMap,
		},
	}
}

var userInput = make(chan []string)

func getInput() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		userInput <- cleanUserInput(scanner.Text())
	}
}

func cleanUserInput(text string) []string {
	output := strings.ToLower(text)
	return strings.Fields(output)
}

func startCli() {

	go getInput()

	fmt.Println("Welcome to the Pokedex! \nUsage:")
	commandhelp()
	fmt.Print("pokedex > ")

	for input := range userInput {
		fmt.Print("pokedex > ")
		command, exists := getCommands()[input[0]]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command please refer to the help guide")
			continue
		}
	}

}
