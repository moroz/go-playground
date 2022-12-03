package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid number of arguments")
		os.Exit(1)
	}

	input := os.Args[1]

	file, err := os.Open(input)
	if err != nil {
		fmt.Printf("Failed to open file %s\n", input)
		os.Exit(1)
	}

	visits := make(map[string]int)
	uniqueVisits := make(map[string]map[string]bool)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Fields(line)
		if len(split) != 2 {
			continue
		}

		url := split[0]
		ip := split[1]

		visits[url]++

		if uniqueVisits[url] == nil {
			uniqueVisits[url] = make(map[string]bool)
		}
		uniqueVisits[url][ip] = true
	}

	for key := range visits {
		fmt.Println(key)
	}

	if scanner.Err() != nil {
		log.Println(scanner.Err())
		os.Exit(1)
	}
}
