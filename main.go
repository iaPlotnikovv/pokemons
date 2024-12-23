package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//fmt.Println("Hello World!")
	scanner := bufio.NewScanner(os.Stdin)
	cmdList := cmdInit()
	for {
		fmt.Print("\n Pokemon >")
		if scanner.Scan() {

			input := CleanInput(scanner.Text())

			if len(input) > 1 {
				fmt.Println("\nWrite a one-word command!")
				continue
			}

			if key, ok := cmdList.commands[input[0]]; ok {

				if err := key.callback(); err != nil {
					fmt.Printf("\nerror!!!: %v\n", err)
				}
				//fmt.Printf("\nYour command was: %v\n", input[0])
			} else {
				fmt.Printf("\nUnknown command: '%v', use 'help'\n", input[0])
			}
		}

	}
}
