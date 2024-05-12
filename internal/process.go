package internal

import (
	"golang.org/x/net/html"
	"strings"
)

var codeTags = map[string]bool{
	"code":     true,
	"pre":      true,
	"textarea": true,
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

func findCodeCells(n *html.Node) []string {
	var codeCells []string

	if n.Type == html.ElementNode && codeTags[n.Data] {
		codeCells = append(codeCells, extractText(n))
	}

	// Traverse child nodes recursively
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		codeCells = append(codeCells, findCodeCells(c)...)
	}

	return codeCells
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
