package main

import "strings"

func CleanInput(text string) []string {
	//The purpose of this function is to split the users input into "words" based on whitespace.
	//It should also lowercase the input and trim any leading or trailing whitespace.
	result := strings.Fields(strings.ToLower(text))

	return result
}
