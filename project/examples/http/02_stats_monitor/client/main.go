package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxDelay = 500
)

func main() {
	req := flag.Int("req", 20, "number of concurrent requests")
	flag.Parse()
	// Configuration
	baseURL := "http://localhost:8080"
	concurrentRequests := *req

	fmt.Printf("Sending %d concurrent requests to %s\n", concurrentRequests, baseURL)
	fmt.Println("==============================================")

	// Create wait group to track all goroutines
	var wg sync.WaitGroup
	wg.Add(concurrentRequests)

	startTime := time.Now()

	successCount := atomic.Int32{}

	for i := range concurrentRequests {
		go func(requestNum int, successCount *atomic.Int32) {
			time.Sleep(time.Duration(rand.Intn(maxDelay)) * time.Millisecond)
			defer wg.Done()
			start := time.Now()
			resp, err := http.Get(baseURL)
			if err != nil {
				fmt.Printf("Request %d failed: %v\n", requestNum, err)
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Error reading response %d: %v\n", requestNum, err)
				return
			}

			elapsed := time.Since(start)
			successCount.Add(1)
			fmt.Printf("Request %d completed in %v\n", requestNum, elapsed)
			fmt.Printf("Response: %s\n", string(body))
		}(i+1, &successCount)
	}

	// Wait for all requests to complete
	wg.Wait()

	// Get final stats
	statsResp, err := http.Get(baseURL + "/stats")
	if err == nil {
		defer statsResp.Body.Close()
		statsBody, _ := io.ReadAll(statsResp.Body)
		fmt.Println("\nFinal server stats:")
		fmt.Println(string(statsBody))
	}

	// Print summary
	totalTime := time.Since(startTime)
	fmt.Printf("\n%d requests completed in %v\n", successCount.Load(), totalTime)
	fmt.Printf("\nFailed requests: %d\n", concurrentRequests-int(successCount.Load()))
	fmt.Printf("\nAverage time per request: %v\n", totalTime/time.Duration(successCount.Load()))
}
