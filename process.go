package main

import (
	"golang.org/x/net/html"
	"strings"
)

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
