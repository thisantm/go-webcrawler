package main

import (
	"fmt"
	"sort"
)

type PageCount struct {
	URL   string
	Count int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`
=============================
REPORT for %s
=============================
`, baseURL)

	pageCounts := sortPages(pages)

	for _, pageCount := range pageCounts {
		fmt.Printf("Found %d internal links to %s\n", pageCount.Count, pageCount.URL)
	}
}

func sortPages(pages map[string]int) []PageCount {
	pageSlice := make([]PageCount, 0, len(pages))

	for url, count := range pages {
		pageSlice = append(pageSlice, PageCount{URL: url, Count: count})
	}

	sort.Slice(pageSlice, func(i, j int) bool {
		if pageSlice[i].Count == pageSlice[j].Count {
			return pageSlice[i].URL < pageSlice[j].URL
		}
		return pageSlice[i].Count > pageSlice[j].Count
	})

	return pageSlice
}
