package wordcount

import (
	"bufio"
	"os"
	"strings"
)

// CountWordsInFile counts the words in a file
func CountWordsInFile(filename string) (int, error) {
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
