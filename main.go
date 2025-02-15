package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// Task represents a unit of work
type Task struct {
	ID int
}

// Worker processes tasks with controlled concurrency using semaphore
func worker(id int, tasks <-chan Task, wg *sync.WaitGroup, sem *semaphore.Weighted) {
	defer wg.Done()

	for task := range tasks {
		// Acquire semaphore permit
		err := sem.Acquire(context.Background(), 1)
		if err != nil {
			fmt.Printf("Failed to acquire semaphore: %v\n", err)
			return
		}

		fmt.Printf("Worker %d starting task %d\n", id, task.ID)
		// Simulate work with random duration
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Printf("Worker %d completed task %d\n", id, task.ID)

		// Release semaphore permit
		sem.Release(1)
	}
}

func main() {
	// Set random seed
	rand.Seed(time.Now().UnixNano())

	// Create a buffered channel for tasks
	taskQueue := make(chan Task, 10)
	
	// Create a weighted semaphore with max 3 concurrent operations
	sem := semaphore.NewWeighted(3)
	
	var wg sync.WaitGroup
	
	// Create 5 workers
	numWorkers := 5
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, taskQueue, &wg, sem)
	}

	// Send 10 tasks to the workers
	numTasks := 10
	fmt.Printf("Dispatching %d tasks to %d workers (max 3 concurrent tasks)...\n", numTasks, numWorkers)
	for i := 1; i <= numTasks; i++ {
		taskQueue <- Task{ID: i}
	}
	close(taskQueue)

	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("All tasks completed")
}
