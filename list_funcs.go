package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/iaPlotnikovv/pokemons/internal/pokeapi"
	cache "github.com/iaPlotnikovv/pokemons/internal/pokecache"
)

type CommandCLI struct {
	Name     string
	desc     string
	callback func(*pokeapi.Config, *cache.Cache) error
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
	cmd.commands["cache"] = CommandCLI{
		Name:     "cache",
		desc:     "Show cache",
		callback: cacheCheck,
	}

	return cmd
}

func CleanInput(text string) []string {
	//The purpose of this function is to split the users input into "words" based on whitespace.
	//It should also lowercase the input and trim any leading or trailing whitespace.
	result := strings.Fields(strings.ToLower(text))

	return result
}

func commandExit(c *pokeapi.Config, cache *cache.Cache) error {

	fmt.Printf("\nClosing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func (c *CommandsList) commandHelp(p *pokeapi.Config, cache *cache.Cache) error {

	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, v := range c.commands {
		fmt.Printf("%v: %v\n", v.Name, v.desc)
	}
	fmt.Printf("\nYOU'RE ON PAGE: %v\n\n", p.Page)
	return nil
}
func CommandMap(url *pokeapi.Config, cache *cache.Cache) error {
	start := time.Now()
	var apiResp pokeapi.APIResponse
	//checking cache
	fmt.Printf("\nChecking cache for URL: %s\n", url.Next)
	if data, ok := cache.Get(url.Next); ok {

		if err := json.Unmarshal(data, &apiResp); err != nil {
			return err
		}
		for _, loc := range apiResp.Results {
			fmt.Println(loc.Name)
		}

	} else {
		// request for data
		res, err := http.Get(url.Next)
		if err != nil {
			fmt.Printf("\nerror in get!!: %v", err)
			return err
		}
		defer res.Body.Close()

		decoder := json.NewDecoder(res.Body)

		// storing data in cache

		if err = decoder.Decode(&apiResp); err != nil {
			return err
		}

		cachedData, err := json.Marshal(apiResp)
		if err != nil {
			return err
		}
		cache.Add(url.Next, cachedData)

		for _, loc := range apiResp.Results {
			fmt.Println(loc.Name)
		}

	}
	fmt.Printf("\nYOU'RE ON PAGE: %v\n\n", url.Page)
	url.Update(apiResp.Next, apiResp.Previous)
	fmt.Println(time.Since(start))

	return nil

}

func CommandMapBack(url *pokeapi.Config, cache *cache.Cache) error {
	if url.Previous != nil {
		url.Next, url.Previous = *url.Previous, &url.Next
		err := CommandMap(url, cache)
		return err
	}
	fmt.Print("\nu on the first bruh\n")
	return nil
}

func cacheCheck(url *pokeapi.Config, cache *cache.Cache) error {
	if len(cache.Data) != 0 {
		for k := range cache.Data {
			fmt.Println(k)
		}
	} else {
		fmt.Print("\nCache is empty!\n")
	}

	return nil
}
