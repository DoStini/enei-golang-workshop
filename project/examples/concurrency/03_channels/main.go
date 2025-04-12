package main

import (
	"flag"
	"fmt"
	"sync"
)

type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value++
}
func (c *Counter) GetValue() int {
	return c.value
}

// runWithRaceCondition demonstrates a race condition when multiple
// goroutines access shared memory without synchronization
func runWithRaceCondition() {
	fmt.Println("\n=== Running example with race condition ===")
	counter := Counter{value: 0}

	// Number of increments to perform
	const numIncrements = 1000
	const numGoroutines = 1000
	var wg sync.WaitGroup

	// Launch multiple goroutines to increment the counter
	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Each goroutine increments the counter multiple times
			for range numIncrements {
				// RACE CONDITION: Multiple goroutines access and modify
				// the same memory location without synchronization
				counter.Increment()
			}
		}()
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// The expected value is 10 goroutines × 1000 increments = 10000
	// But due to race conditions, we often get a lower value
	fmt.Printf("\nExpected counter value: %d\n", numGoroutines*numIncrements)
	fmt.Printf("\nActual counter value: %d\n", counter.GetValue())
	fmt.Println("\nThe actual value may be less than expected due to race conditions")
}

// runWithChannels demonstrates how to fix race conditions using channels
// for communication instead of directly accessing shared memory
func runWithChannels() {
	fmt.Println("=== Running example with channels ===")

	const numIncrements = 1000
	const numGoroutines = 1000

	// Create a channel to receive increment commands
	// This channel follows the "share memory by communicating" pattern
	incrementCh := make(chan struct{})

	// Create a channel to signal when we're done sending commands
	doneCh := make(chan struct{})

	// Create a channel to get the final counter value
	resultCh := make(chan int)

	// Launch a single goroutine to manage the counter
	// This eliminates the race condition by having only one goroutine
	// access the counter value
	go func() {
		counter := 0

		for {
			select {
			case <-incrementCh:
				// Increment the counter when requested
				counter++
			case <-doneCh:
				// When done, send the final value and exit
				resultCh <- counter
				return
			}
		}
	}()

	// Launch multiple goroutines that send increment requests
	var wg sync.WaitGroup
	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Send increment requests through the channel
			for range numIncrements {
				incrementCh <- struct{}{}
			}
		}()
	}

	// Wait for all increment requests to be sent
	wg.Wait()

	// Signal that we're done sending increment requests
	doneCh <- struct{}{}

	// Get and display the final result
	finalCount := <-resultCh
	fmt.Printf("\nExpected counter value: %d\n", numGoroutines*numIncrements)
	fmt.Printf("\nActual counter value: %d\n", finalCount)
	fmt.Println("\nUsing channels ensures the correct result")
}

func main() {
	// Command line flag to choose which example to run
	channels := flag.Bool("chan", false, "Use channels to avoid race conditions")
	flag.Parse()
	shouldUseChannels := *channels

	if shouldUseChannels {
		runWithChannels()
	} else {
		runWithRaceCondition()
	}
}
