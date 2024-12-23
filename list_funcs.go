package main

import (
	"fmt"
	"os"
	"strings"
)

type commandCLI struct {
	name     string
	desc     string
	callback func() error
}

type CommandsList struct {
	commands map[string]commandCLI
}

func cmdInit() *CommandsList {
	cmd := &CommandsList{
		commands: make(map[string]commandCLI),
	}

	cmd.commands["help"] = commandCLI{
		name:     "help",
		desc:     "Displays a help message",
		callback: cmd.commandHelp,
	}
	cmd.commands["exit"] = commandCLI{
		name:     "exit",
		desc:     "Exit the Pokedex",
		callback: commandExit,
	}

	return cmd
}

func CleanInput(text string) []string {
	//The purpose of this function is to split the users input into "words" based on whitespace.
	//It should also lowercase the input and trim any leading or trailing whitespace.
	result := strings.Fields(strings.ToLower(text))

	return result
}

func commandExit() error {

	fmt.Printf("\nClosing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func (c *CommandsList) commandHelp() error {

	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, v := range c.commands {
		fmt.Printf("%v: %v\n", v.name, v.desc)
	}
	return nil
}
