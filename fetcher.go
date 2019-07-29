package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strings"
)

func makeRequest(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func fetchLinks(doc *goquery.Document) []string {
	links := []string{}
	if doc != nil {
		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			resp, _ := s.Attr("href")
			links = append(links, resp)
		})
		return links
	}
	return links
}

func resolveRelativeLinks(baseurl string, hrefs []string) []string {
	urls := []string{}
	for _, href := range hrefs {
		// If the url starts with baseurl add it to list
		if strings.HasPrefix(href, baseurl) {
			urls = append(urls, href)
		}
		// If the url is a relative url, add baseurl to create a proper url
		if strings.HasPrefix(href, "/") {
			url := fmt.Sprintf("%s%s", baseurl, href)
			urls = append(urls, url)
		}
	}
	return urls
}

func getBaseURL(u string) (baseurl string, err error) {
	parsed, err := url.Parse(u)
	if err != nil {
		fmt.Printf("Unable to get baseurl of %v\n", u)
		return "", err
	}
	return fmt.Sprintf("%s://%s", parsed.Scheme, parsed.Host), nil
}

// Fetch method will fetch the links from a single page
func (r Result) Fetch(url string) (title string, links []string, err error) {
	resp, err := makeRequest(url)
	if err != nil {
		fmt.Printf("Failed to GET the page %v\n", url)
		return "", nil, err
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return "", nil, err
	}
	title = doc.Find("title").First().Text()
	links = fetchLinks(doc)
	baseurl, err := getBaseURL(url)
	if err != nil {
		return "", nil, err
	}
	links = resolveRelativeLinks(baseurl, links)
	return title, links, nil
}
