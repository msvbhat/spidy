package main

import (
	"fmt"
)

func main() {
	url := "https://google.com"
	result := Result{}
	title, links, err := result.Fetch(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(title)
	fmt.Println(links)
}
