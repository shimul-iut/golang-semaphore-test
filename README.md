# Go Semaphore Example

This project demonstrates the implementation of a weighted semaphore in Go using the `golang.org/x/sync/semaphore` package. The example shows how to control concurrent access to resources using semaphores in a worker pool pattern.

## Features

- Weighted semaphore implementation
- Worker pool pattern with 5 workers
- Concurrent task processing with controlled access (max 3 concurrent tasks)
- Random task duration simulation
- Channel-based task distribution

## Implementation Details

The project implements:
- A semaphore that limits concurrent execution to 3 tasks
- 5 workers competing for the semaphore permits
- 10 tasks distributed through a buffered channel
- Random processing time for tasks to demonstrate concurrent execution

### Key Components

- `semaphore.NewWeighted(3)`: Creates a semaphore allowing max 3 concurrent operations
- `sem.Acquire()`: Obtains a permit before starting work
- `sem.Release()`: Releases the permit after work completion
- Task struct for representing work units
- Buffered channel for task distribution

## Requirements

- Go 1.21 or later
- golang.org/x/sync package

## Running the Project

1. Clone the repository:
```bash
git clone https://github.com/shimul-iut/golang-semaphore-test.git
```

2. Navigate to the project directory:
```bash
cd golang-semaphore-test
```

3. Run the program:
```bash
go run main.go
```

## Expected Output

You'll see workers processing tasks with a maximum of 3 tasks running concurrently at any time. The output will show:
- Task start and completion times
- Which worker is handling which task
- Concurrent execution of tasks (limited to 3)
