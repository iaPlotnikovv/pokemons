package main

import "fmt"

func pokeInspect(ctx *AppCtx, input []string) error {
	if len(input) == 0 {
		fmt.Printf("\nInvalid command! Use `inspect <pokemon-name>`\n")
		return nil
	}

	if val, ok := ctx.Pokedex.Pokemons[input[0]]; ok {

		ctx.Pokedex.Inspect(&val)
		//fmt.Println(val)
		return nil

	} else {
		fmt.Printf("\nYou haven't got this pokemon in Pokedex!\n")
		return nil
	}

}
