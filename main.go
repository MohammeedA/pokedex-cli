package main

import (
	"strings"
)

// cleanInput takes a string input and returns a slice of strings where:
// - The resulting string is split into words
// - Leading and trailing whitespace is removed
// - All text is converted to lowercase
//
// Example:
//
//	cleanInput("  Hello   WORLD  ") returns []string{"hello", "world"}
//	cleanInput("testnowhitespace") returns []string{"testnowhitespace"}
func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	return words
}

func main() {
	// This is a simple Go program that prints "Hello, World!" to the console.
	println("Hello, World!")
}
