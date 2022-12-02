package main

import "fmt"

func main() {
	// Use hard-coded text
	text := "let's count some words!"

	var numSpaces int

	// Scan the text for spaces
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			numSpaces++
		}
	}

	fmt.Println("Found", numSpaces, "spaces, so there must be", numSpaces+1, "words")
}
