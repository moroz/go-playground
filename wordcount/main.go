package main

import (
	"fmt"
	"strings"
)

func main() {
	// Use hard-coded text
	text := "let's count some words!"
	words := strings.Fields(text)
	fmt.Println("Found", len(words), "words")
}
