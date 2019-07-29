package main

import (
	"fmt"
)

func main() {
	url := "https://monzo.com"
	res := Result{}
	Crawl(url, res)
	fmt.Printf("The links found in %s are...\n", url)
	for u, _ := range result.cmap {
		fmt.Println(u)
	}
}
