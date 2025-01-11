package main

import (
	"fmt"
	"os"
	"time"

	"math/rand"

	"github.com/iaPlotnikovv/pokemons/internal/pokeapi"
)

func pokeExplore(ctx *AppCtx, input []string) error {
	start := time.Now()
	if len(input) == 0 {
		fmt.Printf("Invalid command! Use `explore <area-field>`")
		return nil
	}

	//fmt.Println("\nHERE!!!\n", pointer[k].Url)
	areaUrl := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", input[0])

	apiResp, err := pokeapi.GetReq(areaUrl, ctx.Cache)
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

func commandExit(ctx *AppCtx, input []string) error {
	if len(input) != 0 {
		fmt.Printf("\n\nToo many arguments for ONE-WORD command!\n\n")
		return nil
	}
	fmt.Printf("\nClosing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
func CommandMap(ctx *AppCtx, input []string) error {
	if len(input) != 0 {
		fmt.Printf("\n\nToo many arguments for ONE-WORD command!\n\n")
		return nil
	}
	start := time.Now()

	apiResp, err := pokeapi.GetReq(ctx.Config.Next, ctx.Cache)
	if err != nil {
		fmt.Printf("\nerror in get!!: %v", err)
	}

	for _, loc := range apiResp.Results {
		fmt.Println(loc.Name)
	}

	fmt.Printf("\nYOU'RE ON PAGE: %v\n\n", ctx.Config.Page)
	ctx.Config.Update(apiResp.Next, apiResp.Previous)
	//url.TakeResult(&apiResp)
	fmt.Println(time.Since(start))
	return nil

}

func CommandMapBack(ctx *AppCtx, input []string) error {
	if len(input) != 0 {
		fmt.Printf("\n\nToo many arguments for ONE-WORD command!\n\n")
		return nil
	}

	if ctx.Config.Previous != nil {
		ctx.Config.Next, ctx.Config.Previous = *ctx.Config.Previous, &ctx.Config.Next
		err := CommandMap(ctx, input)
		return err
	}
	if ctx.Config.Page == 0 {
		fmt.Printf("\n\nTry `map` to open pages!\n\n")
		return nil
	} else if ctx.Config.Page == 1 {
		fmt.Print("\nu on the first bruh\n")

	}

	return nil
}

func cacheCheck(ctx *AppCtx, input []string) error {
	if len(input) != 0 {
		fmt.Printf("\n\nToo many arguments for ONE-WORD command!\n\n")
	}
	if len(ctx.Cache.Data) != 0 {
		for k := range ctx.Cache.Data {
			fmt.Println(k)
		}
	} else {
		fmt.Print("\nCache is empty!\n")
	}

	return nil
}

func pokeCatch(ctx *AppCtx, input []string) error {

	if len(input) == 0 {
		fmt.Printf("Invalid command! Use `catch <pokemon-name>`")
		return nil
	}
	if _, ok := ctx.Pokedex.Pokemons[input[0]]; ok {
		fmt.Printf("This pokemon is already in Pokedex!")
		return nil
	}
	poke, err := ctx.Pokedex.Add(input[0])
	if err != nil {
		return err
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	roll := r.Intn(20) + 1
	fmt.Printf("\n\nYour dice roll is %v\n\n", roll)

	time.Sleep(1 * time.Second)
	fmt.Printf("\n\nThrowing a Pokeball at %s...\n\n", poke.Name)
	time.Sleep(1 * time.Second)

	catch := ctx.Pokedex.RollCheck(roll, &poke)
	if catch {
		fmt.Printf("\n\n%s was caught!\n\n", poke.Name)
		ctx.Pokedex.Pokemons[poke.Name] = poke //pokemon added to pokedex!
		return nil
	} else {
		fmt.Printf("\n\n%s escaped!\n\n", poke.Name)
		return nil
	}
}

func pokeDex(ctx *AppCtx, input []string) error {
	if len(ctx.Pokedex.Pokemons) != 0 {
		for k := range ctx.Pokedex.Pokemons {
			fmt.Printf("\n%s\n", k)
		}
	} else {
		fmt.Printf("\nYour Pokedex is empty! Try to catch a pokemon!\n")
	}

	return nil
}
