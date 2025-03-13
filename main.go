package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
	"time"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func main() {
	t := time.Now()
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL, err := url.Parse(args[0])
	if err != nil {
		fmt.Printf("failed to parse base URL: %v", err)
		os.Exit(1)
	}

	cfg := config{
		pages:              map[string]int{},
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, 5),
		wg:                 &sync.WaitGroup{},
	}

	fmt.Printf("starting crawl of: %s\n", cfg.baseURL.String())
	cfg.crawlPage(cfg.baseURL.String())
	cfg.wg.Wait()

	fmt.Println("\n================ REPORT ================")
	for k, v := range cfg.pages {
		fmt.Printf("%v: %v\n", k, v)
	}

	fmt.Printf("\ntime to run: %v\n", time.Since(t))
}
