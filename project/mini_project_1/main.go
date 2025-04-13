package main

import (
	"fmt"
	"net/http"
)

func CheckUrls(urls []string) map[string]bool {
	return map[string]bool{}
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
	}

	// Task: Define a results map here

	// Task: Loop through the URL's and gather their response
	resp, _ := http.Get(urls[0])
	// Task: Error handling
	// Task: Defer the closure of the resp.Body
	// Task: Save in the map true if 200, false otherwise

	fmt.Println(resp.StatusCode)

	// Task: Print the result map

	// Task: Print the result map with formatting +v

	// Task: iterate over the map printing key and value with formatting
}
