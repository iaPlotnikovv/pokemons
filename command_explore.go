package main

import (
	"fmt"
	"time"

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
	ctx.Pokearea.Refresh()
	for _, p := range apiResp.Pokemons {
		fmt.Println(p.Pokemon.Name)
		ctx.Pokearea.Update(p.Pokemon.Name)
	}

	fmt.Println("\n", time.Since(start))
	return nil
}
