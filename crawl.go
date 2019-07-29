package main

// Crawl Function/Method to recursively crawl the urL
func (result *Result) Crawl(url string, fetcher Fetcher) {
	// This Lock is to avoid processing the same link/url twice
	result.Lock()
	if _, ok := result.cmap[url]; ok {
		result.Unlock()
		return
	}
	// Add the url to the map while the lock is being held
	result.cmap[url] = "processing"
	result.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		return
	}
	// Update the result after acquiring the Lock
	result.Lock()
	result.cmap[url] = body
	//result.cmap[url] = append(result.cmap[url], urls...)
	result.Unlock()

	// A channel to block till all links are processed
	done := make(chan bool)
	for _, u := range urls {
		go func(url string) {
			result.Crawl(url, fetcher)
			done <- true
		}(u)
	}

	// Now wait for those results from for loop to be ready
	for range urls {
		<-done
	}
}
