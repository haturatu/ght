package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/html"
)

func findTitleTag(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "title" {
		if n.FirstChild != nil {
			return n.FirstChild.Data
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result := findTitleTag(c); result != "" {
			return result
		}
	}
	return ""
}

func fetchAndParse(client *http.Client, url string, useRange bool) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	if useRange {
		req.Header.Set("Range", "bytes=0-4096")
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %w", err)
	}

	return findTitleTag(doc), nil
}

func fetchTitle(url string) (string, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// range limit : get request
	title, err := fetchAndParse(client, url, true)
	if err != nil {
		return "", err
	}
	if title != "" {
		return title, nil
	}

	// no range limit : get reqest
	title, err = fetchAndParse(client, url, false)
	if err != nil {
		return "", err
	}
	if title == "" {
		return "", fmt.Errorf("no title found: %s", url)
	}

	return title, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: ght \"https://google.com/\"")
		os.Exit(1)
	}

	url := os.Args[1]
	title, err := fetchTitle(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(2)
	}

	fmt.Printf("%s\n", title)
}
