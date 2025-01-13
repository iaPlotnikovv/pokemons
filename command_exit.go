package main

import (
	"fmt"
	"os"
)

func commandExit(ctx *AppCtx, input []string) error {
	if len(input) != 0 {
		fmt.Printf("\n\nToo many arguments for ONE-WORD command!\n\n")
		return nil
	}
	fmt.Printf("\nClosing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
