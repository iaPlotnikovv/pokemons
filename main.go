package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/iaPlotnikovv/pokemons/internal/pokeapi"
	cache "github.com/iaPlotnikovv/pokemons/internal/pokecache"
)

func main() {
	//fmt.Println("Hello World!")
	scanner := bufio.NewScanner(os.Stdin)
	cmdList := cmdInit()
	configL := pokeapi.InitConf()
	caching := cache.NewCache(20 * time.Second)

	//Welcome msg
	cmdList.commandHelp(configL, caching, []string{})
	//Welcome msg

	for {
		fmt.Print("\n Pokemon >")
		if scanner.Scan() {

			input := CleanInput(scanner.Text())

			//if len(input) > 1 || len(input) == 0 {
			//	fmt.Println("\nWrite a one-word command!")
			//	continue
			//}

			if key, ok := cmdList.commands[input[0]]; ok {

				configL.PageCounter(input)

				if err := key.callback(configL, caching, input[1:]); err != nil {
					fmt.Printf("\nerror!!!: %v\n", err)
				}
				//fmt.Printf("\nYour command was: %v\n", input[0])
			} else {
				fmt.Printf("\nUnknown command: '%v', use 'help'\n", input[0])
			}
		}

	}
}
