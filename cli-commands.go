package main

import (
	"fmt"
	"os"
)

func commandhelp() error {
	for _, cmd := range getCommands() {
		fmt.Printf("\n%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandMap() error {
	//TODO code for this func
	return nil
}
