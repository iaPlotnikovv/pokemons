package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//fmt.Println("Hello World!")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n Pokemon >")
		if scanner.Scan() {

			input := CleanInput(scanner.Text())

			fmt.Printf("\nYour command was: %v\n", input[0])

		}
	}

}
