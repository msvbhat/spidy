package main

type Fetcher interface {
	Fetch(url string) (title string, urls []string, err error)
}

var result = Result{cmap: make(map[string][]string)}

func Crawl(url string, fetcher Fetcher) {
	// Acquire the lock and check if the url is already present in the map.
	// This is for not processing the same link more than once.
	result.Lock()
	if _, ok := result.cmap[url]; ok {
		result.Unlock()
		return
	}
	// Add the url to the map while the lock is being held
	result.cmap[url] = []string{}
	result.Unlock()

	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		return
	}
	// Update the result after acquiring the Lock
	result.Lock()
	result.cmap[url] = append(result.cmap[url], urls...)
	result.Unlock()

	// A channel to block till all links are processed
	done := make(chan bool)
	for _, u := range urls {
		go func(url string) {
			Crawl(url, fetcher)
			done <- true
		}(u)
	}

	// Now wait for those results from for loop to be ready
	for _ = range urls {
		<-done
	}
}
