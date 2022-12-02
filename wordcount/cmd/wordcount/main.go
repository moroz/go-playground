package main

import (
	"fmt"
	"gobook/wordcount"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("need to provide filename!")
		os.Exit(1)
	}

	wordCount, err := wordcount.CountWordsInFile(os.Args[1])

	// This will be not nil if scanner.Scan() returned early
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println("Found", wordCount, "words")
}
