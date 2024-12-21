package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"ght/chardet"

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

	// encoding and decode
	body, err := chardet.DetectAndDecode(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to decode response body: %w", err)
	}

	doc, err := html.Parse(body)
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

	// no range limit : get request
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
	markdown := flag.Bool("m", false, "Output the URL in Markdown format")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <URL>\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	url := flag.Arg(0)
	title, err := fetchTitle(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(2)
	}

	// output md
	if *markdown {
		fmt.Printf("[%s](%s)\n", title, url)
	} else {
		fmt.Printf("%s\n", title)
	}
}
