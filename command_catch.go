package main

import (
	"fmt"
	"time"

	"math/rand"
)

func pokeCatch(ctx *AppCtx, input []string) error {

	if len(input) == 0 {
		fmt.Printf("Invalid command! Use `catch <pokemon-name>`")
		return nil
	}

	s := CheckName(ctx, input[0])
	if len(s) != 0 {
		fmt.Println(s)
		return nil
	}

	poke, err := ctx.Pokedex.Add(input[0])
	if err != nil {
		return err
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	roll := r.Intn(20) + 1
	//roll := 1
	fmt.Printf("\n\n- Pokemon's base experience is %v\n\n- Your dice roll is %v\n\n", poke.Base_experience, roll)

	time.Sleep(1 * time.Second)
	fmt.Printf("\n\nThrowing a Pokeball at %s...\n\n", poke.Name)
	time.Sleep(1 * time.Second)

	catch := ctx.Pokedex.RollCheck(roll, &poke)
	if catch {
		fmt.Printf("\n\n%s was caught!\n\n", poke.Name)
		ctx.Pokedex.Pokemons[poke.Name] = poke //pokemon added to pokedex!
		fmt.Printf("\nYou may now inspect it with the inspect command\n")
		return nil
	} else {
		fmt.Printf("\n\n%s escaped!\n\n", poke.Name)
		ctx.Ban.Fail(poke.Name)
		return nil
	}
}
