package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func iterateOverUrlSync(urls []string) map[string]bool {
	result := make(map[string]bool)
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		result[url] = resp.StatusCode == 200
	}
	return result
}

func iterateOverUrlAsync(urls []string) map[string]bool {
	initialTime := time.Now()
	var mu sync.Mutex
	result := make(map[string]bool)
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()
			start := time.Now()
			resp, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			mu.Lock()
			result[url] = resp.StatusCode == 200
			mu.Unlock()
			elapsed := time.Since(start)
			fmt.Printf("Time taken: %s\n", elapsed)
		}(url)
	}

	wg.Wait()
	endTime := time.Since(initialTime)
	fmt.Printf("Time taken: %s\n", endTime)
	return result
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
	}
	result := iterateOverUrlAsync(urls)

	fmt.Printf("%+v\n", result)

	for key, value := range result {
		fmt.Printf("%s: %+v\n", key, value)
	}

}
