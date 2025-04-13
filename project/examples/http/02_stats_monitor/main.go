package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// Global counter to track concurrent requests
var (
	activeRequests int
	maxRequests    int
	requestsMutex  sync.Mutex
)

// Handler function with artificial delay to simulate work
func handler(w http.ResponseWriter, r *http.Request) {
	requestID := rand.Intn(1000)
	startTime := time.Now()

	// Track active requests with mutex for thread safety
	requestsMutex.Lock()
	activeRequests++
	if activeRequests > maxRequests {
		maxRequests = activeRequests
	}
	currentActive := activeRequests
	requestsMutex.Unlock()

	// Log when request starts
	log.Printf("[Request #%d] Started. Active requests: %d", requestID, currentActive)

	// Simulate variable workload (between 50-300 milliseconds)
	processingTime := time.Duration(50+rand.Intn(200)) * time.Millisecond
	time.Sleep(processingTime)

	// Process request info
	requestInfo := fmt.Sprintf("Request #%d processed in %v\n",
		requestID, time.Since(startTime))

	// Write response
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	var res []byte
	fmt.Appendf(res, "Hello, Gopher!\n%s", requestInfo)
	w.Write(res)

	// Update active request count and log completion
	requestsMutex.Lock()
	activeRequests--
	finalActive := activeRequests
	requestsMutex.Unlock()

	log.Printf("[Request #%d] Completed after %v. Active requests: %d",
		requestID, time.Since(startTime), finalActive)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	requestsMutex.Lock()
	stats := fmt.Sprintf("Current active requests: %d\nMax concurrent requests: %d\n",
		activeRequests, maxRequests)
	requestsMutex.Unlock()

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(stats))
}

func main() {
	// Register handlers
	http.HandleFunc("/", handler)
	http.HandleFunc("/stats", statsHandler)

	// Configure and start server
	port := 8080
	log.Printf("Starting server on port %d", port)
	log.Printf("Try running multiple requests with: curl http://localhost:%d/", port)
	log.Printf("Check stats with: curl http://localhost:%d/stats", port)
	log.Printf("For concurrent load testing try: ab -n 100 -c 10 http://localhost:%d/", port)

	// Start the server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
