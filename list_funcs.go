package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/iaPlotnikovv/pokemons/internal/pokeapi"
)

type CommandCLI struct {
	Name     string
	desc     string
	callback func(*pokeapi.Config) error
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
		desc:     "Exit the Pokedex",
		callback: commandExit,
	}
	cmd.commands["map"] = CommandCLI{
		Name:     "map",
		desc:     "Show locations",
		callback: CommandMap,
	}
	cmd.commands["mapb"] = CommandCLI{
		Name:     "mapb",
		desc:     "Show prev locations",
		callback: CommandMapBack,
	}

	return cmd
}

func CleanInput(text string) []string {
	//The purpose of this function is to split the users input into "words" based on whitespace.
	//It should also lowercase the input and trim any leading or trailing whitespace.
	result := strings.Fields(strings.ToLower(text))

	return result
}

func commandExit(c *pokeapi.Config) error {

	fmt.Printf("\nClosing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func (c *CommandsList) commandHelp(p *pokeapi.Config) error {

	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, v := range c.commands {
		fmt.Printf("%v: %v\n", v.Name, v.desc)
	}
	fmt.Printf("\nYOU'RE ON PAGE: %v\n\n", p.Page)
	return nil
}
func CommandMap(url *pokeapi.Config) error {

	res, err := http.Get(url.Next)
	if err != nil {
		fmt.Printf("\nerror in get!!: %v", err)
		return err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	var apiResp pokeapi.APIResponse

	if err = decoder.Decode(&apiResp); err != nil {
		return err
	}

	fmt.Printf("YOU'RE ON PAGE: %v\n\n", url.Page)

	for _, loc := range apiResp.Results {
		fmt.Println(loc.Name)
	}

	url.Update(apiResp.Next, apiResp.Previous)

	return nil

}

func CommandMapBack(url *pokeapi.Config) error {
	if url.Previous != nil {
		url.Next, url.Previous = *url.Previous, &url.Next
		err := CommandMap(url)
		return err
	}
	fmt.Println("u on the first bruh")
	return nil
}
