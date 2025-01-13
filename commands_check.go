package main

import "fmt"

func CheckName(ctx *AppCtx, name string) string {

	if _, ok := ctx.Pokedex.Pokemons[name]; ok {
		return "\nThis pokemon is already in Pokedex!\n"

	}
	if _, exist := ctx.Pokearea.Pokemons_on_area[name]; !exist {

		return "\nPokemon isn't on area!\n"

	}

	if _, banned := ctx.Ban.Banlist[name]; banned {
		return "\nYou've already tried to catch this pokemon! Try again later!\n"

	}
	return ""
}

func cacheCheck(ctx *AppCtx, input []string) error {
	if len(input) != 0 {
		fmt.Printf("\n\nToo many arguments for ONE-WORD command!\n\n")
		return nil
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

func pokeDex(ctx *AppCtx, input []string) error {
	if len(input) != 0 {
		fmt.Printf("\n\nToo many arguments for ONE-WORD command!\n\n")
		return nil
	}

	if len(ctx.Pokedex.Pokemons) != 0 {
		for k := range ctx.Pokedex.Pokemons {
			fmt.Printf("Your Pokedex:\n\n- %s\n", k)
		}
	} else {
		fmt.Printf("\nYour Pokedex is empty! Try to catch a pokemon!\n")
	}

	return nil
}
