package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data map[string]string
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

// runWithChannels demonstrates an inefficient approach using channels
// for a simple shared cache
func runWithChannels() {
	fmt.Println("=== Using channels for cache (inefficient approach) ===")

	// Create channels for different operations
	setCh := make(chan struct {
		key   string
		value string
	})
	getCh := make(chan struct {
		key      string
		resultCh chan string
	})

	// Start a goroutine that manages the cache data
	cache := make(map[string]string)
	go func() {
		for {
			select {
			case setOp := <-setCh:
				// Handle set operation
				cache[setOp.key] = setOp.value
			case getOp := <-getCh:
				// Handle get operation
				value, ok := cache[getOp.key]
				if !ok {
					value = ""
				}
				getOp.resultCh <- value
			}
		}
	}()

	// Function to set a value in the cache using channels
	setWithChannel := func(key, value string) {
		setCh <- struct {
			key   string
			value string
		}{key, value}
	}

	// Function to get a value from the cache using channels
	getWithChannel := func(key string) string {
		resultCh := make(chan string)
		getCh <- struct {
			key      string
			resultCh chan string
		}{key, resultCh}
		return <-resultCh
	}

	// Demonstrate the cache operations
	start := time.Now()

	// Perform operations with the channel-based approach
	var wg sync.WaitGroup
	numGoroutines := 1000
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i%10)
			value := fmt.Sprintf("value%d", i)

			// Set a value
			setWithChannel(key, value)

			// Get a value
			_ = getWithChannel(key)
		}(i)
	}

	wg.Wait()

	fmt.Printf("\nChannel-based approach took: %v\n", time.Since(start))
	fmt.Println("\nProblems with this approach:")
	fmt.Println("\n1. Complex code with multiple channels and message types")
	fmt.Println("\n2. Each read/write requires channel communication")
	fmt.Println("\n3. Each read needs a new response channel")
	fmt.Println("\n4. Inefficient for simple data access patterns")
}

// runWithMutex demonstrates a more appropriate approach using a mutex
// for a simple shared cache
func runWithMutex() {
	fmt.Println("=== Using mutex for cache (efficient approach) ===")

	// Create a cache with mutex protection
	var mutex sync.RWMutex
	cache := make(map[string]string)

	// Function to set a value in the cache using mutex
	setWithMutex := func(key, value string) {
		mutex.Lock()
		defer mutex.Unlock()
		cache[key] = value
	}

	// Function to get a value from the cache using mutex
	getWithMutex := func(key string) string {
		mutex.RLock()
		defer mutex.RUnlock()
		value, ok := cache[key]
		if !ok {
			return ""
		}
		return value
	}

	start := time.Now()

	var wg sync.WaitGroup
	const numGoroutines = 1000
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i%10)
			value := fmt.Sprintf("value%d", i)

			setWithMutex(key, value)
			_ = getWithMutex(key)
		}(i)
	}

	wg.Wait()

	fmt.Printf("\nMutex-based approach took: %v\n", time.Since(start))
	fmt.Println("\nBenefits of this approach:")
	fmt.Println("\n1. Simpler, more direct code")
	fmt.Println("\n2. More efficient for frequent read/write operations")
	fmt.Println("\n3. Uses RWMutex to allow concurrent reads")
	fmt.Println("\n4. Better fits the access pattern of a cache")
}

func main() {
	useMutex := flag.Bool("mutex", false, "Use mutex instead of channels")
	flag.Parse()

	if *useMutex {
		runWithMutex()
	} else {
		runWithChannels()
	}
}
