package main

import (
	"fmt"
)

// This function should filter the words as specified.
func filterWords(words []string) []string {
	// TODO: Only modify this function, no others.
	return words
}

func main() {
	fmt.Println("sup.")
	words := []string{
		"this is a string",
		"!this is another one.",
		"this is a third string!",
		"",
		"t!his string is broken.",
		"!pixels add n00bie doo",
		"alright",
		"word.",
		"sup bruh.",
		"foo bar",
		"what else?",
		"i s t h i s a s t r i n g",
		"okay I think I'm done.",
	}
	
	filtered := filterWords(words)
	fmt.Println(len(filtered))
}
