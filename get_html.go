package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to get response: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return "", fmt.Errorf("response has error code %v", resp.StatusCode)
	}
	if !strings.Contains(resp.Header.Get("content-type"), "text/html") {
		return "", fmt.Errorf("content-type header %v is not text/html", resp.Header.Get("content-type"))
	}

	htmlResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	return string(htmlResp), nil
}
