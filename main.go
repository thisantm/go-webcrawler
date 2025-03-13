package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func main() {
	t := time.Now()
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("usage: ./go-webcrawler URL maxConcurrency maxPages")
		os.Exit(1)
	}

	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL, err := url.Parse(args[0])
	if err != nil {
		fmt.Printf("failed to parse base URL: %v", err)
		os.Exit(1)
	}

	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("failed to parse max concurrency, must be a integer: %v", err)
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("failed to parse max pages, must be a integer: %v", err)
		os.Exit(1)
	}

	cfg := config{
		pages:              map[string]int{},
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	fmt.Printf("starting crawl of: %s\n", cfg.baseURL.String())

	cfg.wg.Add(1)
	cfg.crawlPage(cfg.baseURL.String())
	cfg.wg.Wait()

	printReport(cfg.pages, baseURL.String())

	fmt.Printf("\ntime to run: %v\n", time.Since(t))
}
