package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func countWordsInFile(filename string) (int, error) {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var wordCount int

	// for functions just like while in C
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		wordCount += len(words)
	}

	return wordCount, scanner.Err()
}

func main() {
	if len(os.Args) < 2 {
		log.Println("need to provide filename!")
		os.Exit(1)
	}

	wordCount, err := countWordsInFile(os.Args[1])

	// This will be not nil if scanner.Scan() returned early
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println("Found", wordCount, "words")
}
