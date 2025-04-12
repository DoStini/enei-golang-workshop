package main

import (
	"flag"
	"fmt"
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

func runTasksSequentially() {
	fmt.Println("=== Running tasks sequentially ===")
	start := time.Now()

	// Running tasks sequentially
	simpleTask("Task 1", oneSecond)
	simpleTask("Task 2", twoSeconds)
	simpleTask("Task 3", threeSeconds)

	fmt.Printf("Sequential execution took: %s\n\n", time.Since(start))
}

func runTasksConcurrently() {
	fmt.Println("=== Running tasks with goroutines ===")
	start := time.Now()

	// Using goroutines to run tasks concurrently
	go simpleTask("Task A", oneSecond)
	go simpleTask("Task B", twoSeconds)

	// The goal is for them to ask why the last task is not running in a goroutine
	// Add the last goroutine and see what happens
	// Show that the main goroutine is not blocked by the other goroutines
	simpleTask("Task C", threeSeconds)

	fmt.Printf("Concurrent execution took: %s\n", time.Since(start))
}

func main() {
	concurrent := flag.Bool("concurrent", false, "Run tasks concurrently")
	flag.Parse()
	isConcurrent := *concurrent

	if isConcurrent {
		runTasksConcurrently()
	} else {
		runTasksSequentially()
	}
}
