package main

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetBaseURL(t *testing.T) {
	baseurl, err := getBaseURL("https://example.com/some/path")
	if err != nil {
		t.Log("Unable to get the baseurl the URL")
		t.Fail()
	}
	if baseurl != "https://example.com" {
		t.Log("Incorrect baseurl returned")
		t.Fail()
	}
}

func TestResolveRelativeLinks(t *testing.T) {
	baseurl := "https://example.com"
	links := []string{
		"https://example.com/path",
		"https://github.com/user",
		"https://example.com/link",
	}
	rlinks := resolveRelativeLinks(baseurl, links)
	for _, link := range rlinks {
		if !strings.HasPrefix(link, "https://example.com") {
			t.Logf("%s is not a link from example.com domain", link)
			t.Fail()
		}
	}
}

func TestMakeRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`OK`))
	}))
	defer server.Close()

	_, err := makeRequest(server.URL)
	if err != nil {
		t.Logf("makeRequest returned error: %v", err)
		t.Fail()
	}
}

func TestFetchLinks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		io.WriteString(rw, "<html><body><a href=\"https://test.example.com/path\">testlink</a></body></html>")
	}))
	defer server.Close()

	resp, err := makeRequest(server.URL)
	if err != nil {
		t.Logf("makeRequest returned error: %v", err)
		t.Fail()
	}
	doc, _ := goquery.NewDocumentFromResponse(resp)
	links := fetchLinks(doc)
	if len(links) != 1 {
		t.Log("Number of links returned is incorrect")
		t.Fail()
	}

	for _, link := range links {
		if !strings.HasPrefix(link, "https://test.example.com/") {
			t.Log("fetchLinks returned incorrect links")
			t.Fail()
		}
	}
}
