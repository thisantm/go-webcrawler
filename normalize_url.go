package main

import (
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	parseURL, err := url.Parse(strings.ToLower(inputURL))
	if err != nil {
		return "", err
	}

	normalizedURL := parseURL.Hostname() + parseURL.Path

	return normalizedURL, nil
}
