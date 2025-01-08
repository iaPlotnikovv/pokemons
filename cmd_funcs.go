package main

import (
	"fmt"
	"os"
	"time"

	"github.com/iaPlotnikovv/pokemons/internal/pokeapi"
	cache "github.com/iaPlotnikovv/pokemons/internal/pokecache"
)

func pokeExplore(c *pokeapi.Config, cache *cache.Cache, input []string) error {
	start := time.Now()
	if len(input) == 0 {
		fmt.Printf("Invalid command! Use `explore <area-field>`")
		return nil
	}

	//fmt.Println("\nHERE!!!\n", pointer[k].Url)
	areaUrl := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", input[0])

	apiResp, err := pokeapi.GetReq(areaUrl, cache)
	if err != nil {
		return err

	}

	fmt.Printf("\n\nExploring %v...\n\n", input[0])

	for _, p := range apiResp.Pokemons {
		fmt.Println(p.Pokemon.Name)
	}

	fmt.Println("\n", time.Since(start))
	return nil
}

func commandExit(c *pokeapi.Config, cache *cache.Cache, input []string) error {
	if len(input) != 0 {
		fmt.Printf("\n\nToo many arguments for ONE-WORD command!\n\n")
		return nil
	}
	fmt.Printf("\nClosing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
func CommandMap(url *pokeapi.Config, cache *cache.Cache, input []string) error {
	if len(input) != 0 {
		fmt.Printf("\n\nToo many arguments for ONE-WORD command!\n\n")
		return nil
	}
	start := time.Now()

	apiResp, err := pokeapi.GetReq(url.Next, cache)
	if err != nil {
		fmt.Printf("\nerror in get!!: %v", err)
	}

	for _, loc := range apiResp.Results {
		fmt.Println(loc.Name)
	}

	fmt.Printf("\nYOU'RE ON PAGE: %v\n\n", url.Page)
	url.Update(apiResp.Next, apiResp.Previous)
	//url.TakeResult(&apiResp)
	fmt.Println(time.Since(start))
	return nil

}

func CommandMapBack(url *pokeapi.Config, cache *cache.Cache, input []string) error {
	if len(input) != 0 {
		fmt.Printf("\n\nToo many arguments for ONE-WORD command!\n\n")
		return nil
	}

	if url.Previous != nil {
		url.Next, url.Previous = *url.Previous, &url.Next
		err := CommandMap(url, cache, input)
		return err
	}
	if url.Page == 0 {
		fmt.Printf("\n\nTry `map` to open pages!\n\n")
		return nil
	} else if url.Page == 1 {
		fmt.Print("\nu on the first bruh\n")

	}

	return nil
}

func cacheCheck(url *pokeapi.Config, cache *cache.Cache, input []string) error {
	if len(input) != 0 {
		fmt.Printf("\n\nToo many arguments for ONE-WORD command!\n\n")
	}
	if len(cache.Data) != 0 {
		for k := range cache.Data {
			fmt.Println(k)
		}
	} else {
		fmt.Print("\nCache is empty!\n")
	}

	return nil
}
