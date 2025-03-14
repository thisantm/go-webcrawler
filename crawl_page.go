package main

import (
	"fmt"
	"net/url"
	"strings"
)

func (cfg *config) pagesLen() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages)
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.pagesLen() >= cfg.maxPages {
		return
	}

	currURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to parse base URL: %v\n", err)
		return
	}

	if cfg.baseURL.Hostname() != currURL.Hostname() {
		fmt.Println("page is not part of the base URL domain")
		return
	}

	normalCurrURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to normalize URL: %v\n", err)
		return
	}

	if isFirst := cfg.addPageVisit(normalCurrURL); !isFirst {
		return
	}

	pageHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to get page html: %v\n", err)
		return
	}

	urls, err := getURLsFromHTML(pageHTML, cfg.baseURL)
	if err != nil {
		fmt.Printf("failed to get urls from html: %v\n", err)
		return
	}

	for _, url := range urls {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			continue
		}

		fmt.Printf("crawling into: %v\n", url)
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, ok := cfg.pages[normalizedURL]; ok {
		cfg.pages[normalizedURL]++
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}
