package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// fetchHTML retrieves the HTML content from a given URL
func fetchHTML(url string) (*html.Node, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get URL %s: %w", url, err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code %d", response.StatusCode)
	}

	document, err := html.Parse(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}
	return document, nil
}

// extractText recursively extracts and prints text from the HTML nodes
func extractText(node *html.Node) {
	if node.Type == html.TextNode {
		fmt.Print(node.Data)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		extractText(child)
	}
}

func main() {
	url := flag.String("url", "", "URL of the webpage to scrape")
	flag.Parse()

	if *url == "" {
		fmt.Println("Please provide a URL using the -url flag.")
		os.Exit(1)
	}

	document, err := fetchHTML(*url)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Extracted Text:")
	extractText(document)
}
