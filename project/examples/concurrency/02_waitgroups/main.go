package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

const (
	oneSecond    = 1 * time.Second
	twoSeconds   = 2 * time.Second
	threeSeconds = 3 * time.Second
)

func simpleTask(taskName string, duration time.Duration) {
	fmt.Printf("Starting task: %s\n", taskName)
	time.Sleep(duration)
	fmt.Printf("Completed task: %s\n", taskName)
}

// This version demonstrates the issue when goroutines don't complete
func runTasksWithProblem() {
	fmt.Println("=== Running all tasks with goroutines (problematic) ===")
	start := time.Now()

	// All tasks run as goroutines
	go simpleTask("Task A", oneSecond)
	go simpleTask("Task B", twoSeconds)
	go simpleTask("Task C", threeSeconds)

	// Main goroutine doesn't wait and exits immediately
	// You might not see any output from the tasks as the program exits too quickly
	fmt.Printf("Main function finished after: %s\n", time.Since(start))
	fmt.Println("Notice how tasks might not complete before the program exits!")
}

// This version uses WaitGroup to properly wait for all goroutines
func runTasksWithWaitGroup() {
	fmt.Println("=== Running all tasks with goroutines (using WaitGroup) ===")
	start := time.Now()

	// Create a WaitGroup to track goroutines
	var wg sync.WaitGroup

	// Launch Task A as goroutine
	wg.Add(1) // Increment counter before launching goroutine
	go func() {
		defer wg.Done() // Decrement counter when goroutine completes
		simpleTask("Task A", oneSecond)
	}()

	// Launch Task B as goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		simpleTask("Task B", twoSeconds)
	}()

	// Launch Task C as goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		simpleTask("Task C", threeSeconds)
	}()

	// Block until all goroutines complete
	wg.Wait()

	fmt.Printf("All tasks completed after: %s\n", time.Since(start))
}

func main() {
	wg := flag.Bool("wg", false, "Use WaitGroup to wait for goroutines")
	flag.Parse()
	shouldUseWaitGroup := *wg

	if shouldUseWaitGroup {
		runTasksWithWaitGroup()
	} else {
		runTasksWithProblem()
	}
}
