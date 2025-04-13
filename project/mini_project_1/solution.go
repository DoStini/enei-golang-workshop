package main

import (
	"fmt"
	"log"
	"net/http"
)

func CheckUrlsSol(urls []string) map[string]bool {
	result := make(map[string]bool)
	// Task: Make a loop that will call the urls and save the response in the result map

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			log.Println("error getting url", err)
			result[url] = false
			continue
		}
		defer resp.Body.Close()
		log.Println(resp.StatusCode)

		// Task: Save the response in the result map
		result[url] = resp.StatusCode == 200
	}

	return result
}

func mainSol() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
	}
	result := CheckUrlsSol(urls)

	// Task: Print the result map
	fmt.Println(result)

	// Task: Print the result map with formatting +v
	fmt.Printf("%+v\n", result)

	// Task iterate over the map printing key and value with formatting
	for key, value := range result {
		fmt.Printf("%s: %+v\n", key, value)
	}

}
