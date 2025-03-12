package main

import (
	"fmt"
	"net/url"
	"os"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("failed to parse base URL: %v", err)
		os.Exit(1)
	}

	currURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to parse base URL: %v\n", err)
		return
	}

	if currURL.Hostname() != baseURL.Hostname() {
		fmt.Println("page is not part of the base URL domain")
		return
	}

	normalCurrURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to normalize URL: %v\n", err)
		return
	}

	if _, ok := pages[normalCurrURL]; ok {
		pages[normalCurrURL]++
		return
	}

	pages[normalCurrURL] = 1
	pageHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to get page html: %v\n", err)
	}

	urls, err := getURLsFromHTML(pageHTML, rawBaseURL)
	if err != nil {
		fmt.Printf("failed to get urls from html: %v\n", err)
	}

	for _, url := range urls {
		fmt.Printf("crawling into: %v\n", url)
		crawlPage(rawBaseURL, url, pages)
	}
}
