package main

import (
	"fmt"
	"net/http"
)

// Task 1: Use goroutines and wait groups
func CheckUrls(urls []string) map[string]bool {
	result := make(map[string]bool)

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			result[url] = false
			continue
		}
		defer resp.Body.Close()

		// Task: Save the response in the result map
		result[url] = resp.StatusCode == 200
	}

	return result
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
	}
	result := CheckUrls(urls)

	fmt.Println(result)

	fmt.Printf("%+v\n", result)

	for key, value := range result {
		fmt.Printf("%s: %+v\n", key, value)
	}

}
