# Go Semaphore Example

This project demonstrates the implementation of a weighted semaphore in Go using the `golang.org/x/sync/semaphore` package. The example shows how to control concurrent access to resources using semaphores in a worker pool pattern, with a visual timeline representation of task execution.

## Features

- Weighted semaphore implementation
- Worker pool pattern with 5 workers
- Concurrent task processing with controlled access (max 3 concurrent tasks)
- Random task duration simulation
- Channel-based task distribution
- Visual timeline representation of task execution
- Real-time task tracking with timestamps

## Implementation Details

The project implements:
- A semaphore that limits concurrent execution to 3 tasks
- 5 workers competing for the semaphore permits
- 10 tasks distributed through a buffered channel
- Random processing time for tasks to demonstrate concurrent execution
- ASCII-based visualization of task execution timeline

### Key Components

- `semaphore.NewWeighted(3)`: Creates a semaphore allowing max 3 concurrent operations
- `sem.Acquire()`: Obtains a permit before starting work
- `sem.Release()`: Releases the permit after work completion
- Task struct for representing work units
- Buffered channel for task distribution
- TimelineVisualizer for tracking and displaying task execution

### Visualization

The program includes a visual timeline representation that shows:
```
Task Timeline Visualization:
Legend: [2222] represents Task 2 being processed
Timeline: 20:11:43 -> 20:11:45
----------------------------------------
Worker 1: |[111]   [444]    [777]       |
Worker 2: |   [222]    [555]    [888]   |
Worker 3: |      [333]    [666]    [999]|
Worker 4: |         [444]    [777]      |
Worker 5: |            [555]    [000]   |
----------------------------------------
```

Features of the visualization:
- Real-time timestamps showing exact execution timeline
- Visual representation of task duration using ASCII art
- Clear indication of which worker is handling each task
- Easy-to-see concurrent task execution
- Task start `[` and completion `]` markers
- Task ID displayed during execution duration

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

## Understanding the Output

When you run the program, you'll see:
1. Initial message about task dispatch
2. A visual timeline showing:
   - Each worker's activity on a separate line
   - Task execution represented by `[TaskID]`
   - Concurrent execution of up to 3 tasks
   - Exact timestamps for the entire execution
3. The timeline clearly demonstrates how the semaphore limits concurrent execution to 3 tasks, even with 5 workers available
