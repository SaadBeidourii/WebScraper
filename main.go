package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func main() {

	fmt.Println("Enter the URL of the webpage:")
	var url string
	fmt.Scanln(&url)

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
