package main

import (
	"fmt"
	"time"

	"github.com/iaPlotnikovv/pokemons/internal/pokeapi"
)

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
