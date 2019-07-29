package main

import (
	"fmt"
	"testing"
)

type mockResult struct {
	title string
	urls  []string
}

type mockFetcher map[string]mockResult

func (m mockFetcher) Fetch(url string) (string, []string, error) {
	if resp, ok := m[url]; ok {
		return resp.title, resp.urls, nil
	}
	return "", nil, fmt.Errorf("Unable to fetch url: %s", url)
}

var mockfetcher = mockFetcher{
	"https://example.com": mockResult{
		"Example link",
		[]string{
			"https://example.com/link1",
			"https://example.com/link2",
		},
	},
	"https://example.com/link1": mockResult{
		"Example link one",
		[]string{
			"https://example.com/link3",
			"https://example.com/link2",
		},
	},
	"https://example.com/link3": mockResult{
		"Example link three",
		[]string{
			"https://example.com/link1",
			"https://example.com/link7",
		},
	},
}

func TestCrawl(t *testing.T) {
	mockres := Result{cmap: make(map[string]string)}
	mockres.Crawl("https://example.com", mockfetcher)
	if len(mockres.cmap) != 5 {
		t.Log("Results length not matching")
		t.Fail()
	}
	if res, ok := mockres.cmap["https://example.com"]; ok {
		if res != "Example link" {
			t.Log("Title not matching")
			t.Fail()
		}
	} else {
		t.Log("Expected url not present in the result")
		t.Fail()
	}
	if res, ok := mockres.cmap["https://example.com/link1"]; ok {
		if res != "Example link one" {
			t.Log("Title not matching")
			t.Fail()
		}
	} else {
		t.Log("Expected url not present in the result")
		t.Fail()
	}
	if res, ok := mockres.cmap["https://example.com/link2"]; ok {
		if res == "Example link two" {
			t.Log("Error not matching")
			t.Fail()
		}
	} else {
		t.Log("Expected url not present in the result")
		t.Fail()
	}
	if res, ok := mockres.cmap["https://example.com/link3"]; ok {
		if res != "Example link three" {
			t.Log("Title not matching")
			t.Fail()
		}
	} else {
		t.Log("Expected url not present in the result")
		t.Fail()
	}
	if res, ok := mockres.cmap["https://example.com/link7"]; ok {
		if res != "Unable to fetch url: https://example.com/link7" {
			t.Log("Error not matching")
			t.Fail()
		}
	} else {
		t.Log("Expected url not present in the result")
		t.Fail()
	}
}
