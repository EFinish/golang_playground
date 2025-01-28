package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

type Task struct {
	ID    int
	Value int
}

// Worker processes tasks received from the tasks channel and sends results to the results channel.
func Worker(id int, tasks <-chan Task, results chan<- string) {
	for task := range tasks {
		fmt.Printf("Worker %d processing Task %d\n", id, task.ID)
		results <- fmt.Sprintf("Worker %d finished Task %d (Value: %d)", id, task.ID, task.Value)
	}
}

func generateRandomNumber(Max int, Min int) int {
	rand.Seed(uint64(time.Now().UnixNano()))
	return rand.Intn(Max-Min+1) + Min
}

func main() {
	numWorkers := generateRandomNumber(5, 1)
	numTasks := generateRandomNumber(300, 2)

	// create channels for tasks and results
	tasks := make(chan Task, numTasks)
	results := make(chan string, numTasks)

	// start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		go Worker(i, tasks, results)
	}

	// send tasks to the task channel
	for i := 1; i <= numTasks; i++ {
		tasks <- Task{ID: i, Value: i * 10}
	}
	close(tasks) // close the task channel to signal no more tasks

	// collect results
	for i := 1; i <= numTasks; i++ {
		fmt.Println(<-results)
	}

	fmt.Println("All tasks processed.")
}
