package main

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// Task represents a unit of work
type Task struct {
	ID       int
	StartTime time.Time
	EndTime   time.Time
}

// TaskEvent represents an event in the task timeline
type TaskEvent struct {
	WorkerID  int
	TaskID    int
	IsStart   bool
	Timestamp time.Time
}

// TimelineVisualizer handles the visualization of tasks
type TimelineVisualizer struct {
	events []TaskEvent
	mu     sync.Mutex
}

// NewTimelineVisualizer creates a new timeline visualizer
func NewTimelineVisualizer() *TimelineVisualizer {
	return &TimelineVisualizer{
		events: make([]TaskEvent, 0),
	}
}

// AddEvent adds a new event to the timeline
func (tv *TimelineVisualizer) AddEvent(event TaskEvent) {
	tv.mu.Lock()
	defer tv.mu.Unlock()
	tv.events = append(tv.events, event)
}

// Visualize creates and displays the timeline visualization
func (tv *TimelineVisualizer) Visualize() {
	tv.mu.Lock()
	defer tv.mu.Unlock()

	if len(tv.events) == 0 {
		return
	}

	// Find the start and end times
	startTime := tv.events[0].Timestamp
	endTime := startTime
	for _, event := range tv.events {
		if event.Timestamp.After(endTime) {
			endTime = event.Timestamp
		}
	}

	// Create a map of worker timelines
	workerTimelines := make(map[int][]string)
	timelineDuration := endTime.Sub(startTime)
	timelineWidth := 60 // characters wide

	// Initialize timelines with spaces
	for _, event := range tv.events {
		if _, exists := workerTimelines[event.WorkerID]; !exists {
			workerTimelines[event.WorkerID] = make([]string, timelineWidth)
			for i := range workerTimelines[event.WorkerID] {
				workerTimelines[event.WorkerID][i] = " "
			}
		}
	}

	// Plot events on the timeline
	for i := 0; i < len(tv.events); i++ {
		event := tv.events[i]
		position := int(float64(event.Timestamp.Sub(startTime)) / float64(timelineDuration) * float64(timelineWidth-1))
		
		if event.IsStart {
			workerTimelines[event.WorkerID][position] = "["
			// Find corresponding end event
			for j := i + 1; j < len(tv.events); j++ {
				if tv.events[j].TaskID == event.TaskID && !tv.events[j].IsStart {
					endPosition := int(float64(tv.events[j].Timestamp.Sub(startTime)) / float64(timelineDuration) * float64(timelineWidth-1))
					// Fill the duration with task ID
					for k := position + 1; k < endPosition; k++ {
						workerTimelines[event.WorkerID][k] = fmt.Sprintf("%d", event.TaskID)
					}
					workerTimelines[event.WorkerID][endPosition] = "]"
					break
				}
			}
		}
	}

	// Print the timeline
	fmt.Println("\nTask Timeline Visualization:")
	fmt.Println("Legend: [2222] represents Task 2 being processed")
	fmt.Printf("Timeline: %s -> %s\n", startTime.Format("15:04:05"), endTime.Format("15:04:05"))
	fmt.Println(strings.Repeat("-", timelineWidth+20))

	// Sort workers by ID for consistent output
	for workerID := 1; workerID <= len(workerTimelines); workerID++ {
		timeline := workerTimelines[workerID]
		fmt.Printf("Worker %d: |%s|\n", workerID, strings.Join(timeline, ""))
	}
	fmt.Println(strings.Repeat("-", timelineWidth+20))
}

// Worker processes tasks with controlled concurrency using semaphore
func worker(id int, tasks <-chan Task, wg *sync.WaitGroup, sem *semaphore.Weighted, visualizer *TimelineVisualizer) {
	defer wg.Done()

	for task := range tasks {
		// Acquire semaphore permit
		err := sem.Acquire(context.Background(), 1)
		if err != nil {
			fmt.Printf("Failed to acquire semaphore: %v\n", err)
			return
		}

		startTime := time.Now()
		visualizer.AddEvent(TaskEvent{
			WorkerID:  id,
			TaskID:    task.ID,
			IsStart:   true,
			Timestamp: startTime,
		})

		// Simulate work with random duration
		duration := time.Duration(rand.Intn(1000)) * time.Millisecond
		time.Sleep(duration)

		endTime := time.Now()
		visualizer.AddEvent(TaskEvent{
			WorkerID:  id,
			TaskID:    task.ID,
			IsStart:   false,
			Timestamp: endTime,
		})

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
	
	// Create timeline visualizer
	visualizer := NewTimelineVisualizer()
	
	var wg sync.WaitGroup
	
	// Create 5 workers
	numWorkers := 5
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, taskQueue, &wg, sem, visualizer)
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
	
	// Display the visualization
	visualizer.Visualize()
}
