package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseUrl, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %v", err)
	}

	htmlReader := strings.NewReader(htmlBody)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %v", err)
	}

	urls := []string{}
	urls = traverseNodes(doc, baseUrl, urls)

	return urls, nil
}

func traverseNodes(node *html.Node, baseURL *url.URL, urls []string) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, anchor := range node.Attr {
			if anchor.Key == "href" {
				href, err := url.Parse(anchor.Val)
				if err != nil {
					fmt.Printf("failed to parse href '%v': %v\n", anchor.Val, err)
					continue
				}
				resolvedURL := baseURL.ResolveReference(href)
				urls = append(urls, resolvedURL.String())
			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		urls = traverseNodes(child, baseURL, urls)
	}

	return urls
}
