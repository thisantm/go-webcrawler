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
	}

	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	BASE_URL := args[0]
	fmt.Printf("starting crawl of: %s\n", BASE_URL)
	pages := map[string]int{}
	crawlPage(BASE_URL, BASE_URL, pages)

	fmt.Println("\n================ REPORT ================")
	for k, v := range pages {
		fmt.Printf("%v: %v\n", k, v)
	}
}
