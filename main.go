package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	rawBaseURL := args[0]
	maxConcurrencyString := args[1]
	maxPagesString := args[2]

	maxConcurrency, err := strconv.Atoi(maxConcurrencyString)
	if err != nil {
		fmt.Println("max concurrency should be an int")
	}
	maxPages, err := strconv.Atoi(maxPagesString)
	if err != nil {
		fmt.Println("max pages should be an int")
	}

	config, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("could not initialize config: %v\n", err)
	}

	config.wg.Add(1)
	go config.crawlPage(rawBaseURL)
	config.wg.Wait()

	printReport(config.pages, rawBaseURL)
}
