package main

import (
	"fmt"
	"log"
	"math/rand"
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
	result := make(map[string]bool)
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			result[url] = resp.StatusCode == 200
		}(url)
	}
	wg.Wait()
	return result
}

func iterateSyncWithSleep() {
	for i := range 10 {
		sleeptime := time.Duration(rand.Intn(2000)) * time.Millisecond
		fmt.Printf("\nSleeping for %s | Iteration %d\n", sleeptime, i)
		time.Sleep(sleeptime)
	}
}

func iterateAsyncWithSleepWithoutWaitGroup() {
	// Iterate over a range with a sleep random between 0 and 2 seconds for each iteration
	for i := range 10 {
		go func(i int) {
			sleeptime := time.Duration(rand.Intn(2000)) * time.Millisecond
			fmt.Printf("\nSleeping for %s | Iteration %d\n", sleeptime, i)
			time.Sleep(sleeptime)

		}(i)
	}
}

func iterateAsyncWithSleepWithWaitGroup() {
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sleeptime := time.Duration(rand.Intn(2000)) * time.Millisecond
			fmt.Printf("\nSleeping for %s | Iteration %d\n", sleeptime, i)
			time.Sleep(sleeptime)

		}(i)
	}
	wg.Wait()
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
	}
	result := make(map[string]bool)
	// Task: Make a loop that will call the urls and save the response in the result map
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		// Task: Save the response in the result map
		result[url] = resp.StatusCode == 200
	}
	// Task: Print the result map
	fmt.Println(result)
	// Task: Print the result map with formatting +v
	fmt.Printf("%+v\n", result)

	// Task iterate over the map printing key and value with formatting
	for key, value := range result {
		fmt.Printf("%s: %+v\n", key, value)
	}

}
