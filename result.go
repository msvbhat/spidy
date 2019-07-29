package main

import (
	"sync"
)

// Result struct will hold the final result
type Result struct {
	cmap map[string]string
	sync.Mutex
}

// Fetcher interface should be implemented by the Result struct
type Fetcher interface {
	Fetch(url string) (title string, urls []string, err error)
}
