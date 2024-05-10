package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func main() {
	// Specify the URL of the webpage to scrape
	url := "https://nicegui.io/"

	// Send an HTTP GET request to the URL
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to send GET request: %v\n", err)
		return
	}
	defer response.Body.Close()

	// Check if the response status code is OK (200)
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Failed to get webpage. Status code: %d\n", response.StatusCode)
		return
	}

	// Parse the HTML content of the webpage
	doc, err := html.Parse(response.Body)
	if err != nil {
		fmt.Printf("Failed to parse HTML: %v\n", err)
		return
	}

	title := findTitle(doc)
	if title != "" {
		fmt.Printf("Title: %s\n", title)
	} else {
		fmt.Println("Title not found")
	}
}

func findTitle(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "title" {
		return extractText(n)
	}

	// Traverse child nodes recursively
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title := findTitle(c)
		if title != "" {
			return title
		}
	}

	return ""
}

func extractText(n *html.Node) string {
	var sb strings.Builder

	// Process all text nodes under the given node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			sb.WriteString(c.Data)
		}
	}

	return strings.TrimSpace(sb.String())
}
