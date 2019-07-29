package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a url to crawl through")
		fmt.Println("Usage: spidy <url_to_crawl>")
		os.Exit(1)
	}
	url := os.Args[1]
	result := Result{cmap: make(map[string]string)}
	result.Crawl(url, result)
	fmt.Printf("The links found in %s are...\n", url)
	for u, title := range result.cmap {
		fmt.Printf("%s %s\n", u, title)
	}
}
