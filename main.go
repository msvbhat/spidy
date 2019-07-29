package main

import (
	"fmt"
)

func main() {
	url := "https://xkcd.com"
	result := Result{cmap: make(map[string]string)}
	result.Crawl(url, result)
	fmt.Printf("The links found in %s are...\n", url)
	for u, title := range result.cmap {
		fmt.Printf("%s %s\n", u, title)
	}
}
