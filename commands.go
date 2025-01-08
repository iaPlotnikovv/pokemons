package main

import (
	"fmt"
	"strings"

	"github.com/iaPlotnikovv/pokemons/internal/pokeapi"
	cache "github.com/iaPlotnikovv/pokemons/internal/pokecache"
)

func CleanInput(text string) []string {
	//The purpose of this function is to split the users input into "words" based on whitespace.
	//It should also lowercase the input and trim any leading or trailing whitespace.
	result := strings.Fields(strings.ToLower(text))

	return result
}

type CommandCLI struct {
	Name     string
	desc     string
	callback func(*pokeapi.Config, *cache.Cache, []string) error
}

type CommandsList struct {
	commands map[string]CommandCLI
}

func cmdInit() *CommandsList {
	cmd := &CommandsList{
		commands: make(map[string]CommandCLI),
	}

	cmd.commands["help"] = CommandCLI{
		Name:     "help",
		desc:     "Displays a help message",
		callback: cmd.commandHelp,
	}
	cmd.commands["exit"] = CommandCLI{
		Name:     "exit",
		desc:     "Exits the Pokedex",
		callback: commandExit,
	}
	cmd.commands["map"] = CommandCLI{
		Name:     "map",
		desc:     "Shows locations",
		callback: CommandMap,
	}
	cmd.commands["mapb"] = CommandCLI{
		Name:     "mapb",
		desc:     "Shows prev locations",
		callback: CommandMapBack,
	}
	cmd.commands["cache"] = CommandCLI{
		Name:     "cache",
		desc:     "Shows cache",
		callback: cacheCheck,
	}
	cmd.commands["explore"] = CommandCLI{
		Name:     "explore <area-field>",
		desc:     "Shows pokemons on area",
		callback: pokeExplore,
	}

	return cmd
}

func (c *CommandsList) commandHelp(p *pokeapi.Config, cache *cache.Cache, input []string) error {
	if len(input) != 0 {
		cmd := strings.Join(input, " ")
		if v, ok := c.commands[cmd]; ok {
			fmt.Printf("\n\n%v: %v\n\n", v.Name, v.desc)
			return nil
		} else {
			fmt.Printf("\n\nUnknown command! Use `help`\n\n")
			return nil
		}

	}
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, v := range c.commands {
		fmt.Printf("%v: %v\n", v.Name, v.desc)
	}
	fmt.Printf("\nYOU'RE ON PAGE: %v\n\n", p.Page)
	return nil
}
