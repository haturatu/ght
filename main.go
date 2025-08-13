package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"ght/chardet"
	"github.com/akamensky/argparse"
	"github.com/atotto/clipboard"

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
	// Parse command line arguments
	parser := argparse.NewParser("ght", "Get HTML Title")

	// --url option
	// This option is for specifying the URL to fetch
	urlOpt := parser.String("u", "url", &argparse.Options{
		Help: "URL to fetch",
	})

	// URL positional argument
	// This positional argument is for specifying the URL if not provided with --url option
	// It will not show up in the help message
	// Hidden feature: DISABLEDDESCRIPTIONWILLNOTSHOWUP
	// https://github.com/akamensky/argparse/issues/113
	urlPos := parser.StringPositional(&argparse.Options{
		Help: "DISABLEDDESCRIPTIONWILLNOTSHOWUP",
	})

	markdown := parser.Flag("m", "markdown", &argparse.Options{
		Help: "Output in Markdown format",
	})

	copyClip := parser.Flag("c", "copy", &argparse.Options{
		Help: "Copy to clipboard",
	})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Fprint(os.Stderr, parser.Usage(err))
		os.Exit(1)
	}

	var url string
	if urlOpt != nil && *urlOpt != "" {
		url = *urlOpt
	} else if urlPos != nil && *urlPos != "" {
		url = *urlPos
	} else {
		fmt.Fprint(os.Stderr, "Error: URL is required\n")
		// print like curl
		fmt.Fprintf(os.Stderr, "Example: %s [options...] <url>\n", os.Args[0])
		fmt.Fprint(os.Stderr, parser.Usage(nil))
		os.Exit(1)
	}

	// Validate URL
	title, err := fetchTitle(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(2)
	}

	// output md
	var output string
	if *markdown {
		output = fmt.Sprintf("[%s](%s)", title, url)
	} else {
		output = title
	}

	fmt.Println(output)

	if *copyClip {
		err := clipboard.WriteAll(output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error copying to clipboard: %v\n", err)
			os.Exit(3)
		}
	}
}
