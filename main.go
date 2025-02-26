package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	rawBaseURL := args[0]
	fmt.Printf("starting crawl of: %v...\n", rawBaseURL)

	const maxConcurrency = 5
	config, err := configure(rawBaseURL, maxConcurrency)
	if err != nil {
		fmt.Printf("could not initialize config: %v\n", err)
	}

	config.wg.Add(1)
	go config.crawlPage(rawBaseURL)
	config.wg.Wait()

	for normalizedURL, count := range config.pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}
