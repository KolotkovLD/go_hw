package hw05parallelexecution

import (
	"errors"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	// Initialize channels for tasks and errors.
	taskChannel := make(chan Task, len(tasks))
	errorChannel := make(chan error, m)

	// Count the total number of tasks.
	totalTasks := len(tasks)

	// Channel to signal completion of tasks.
	completionSignal := make(chan struct{})

	// Track how many errors have occurred so far.
	errorCounter := 0

	// Goroutine to handle the processing of jobs.
	go func() {
		for job := range taskChannel {
			go func() {
				err := job()
				if err != nil {
					errorChannel <- err
					errorCounter++
					if errorCounter >= m {
						completionSignal <- struct{}{}
						return
					}
				}
			}()
		}
	}()

	// Close the task channel when all tasks have been sent.
	go func() {
		for i := 0; i < totalTasks; i++ {
			taskChannel <- tasks[i]
		}
		close(taskChannel)
	}()

	// Wait for either all tasks to complete or for the error counter to reach the limit.
	select {
	case <-completionSignal:
		return errors.New("maximum error count exceeded")
	default:
		<-completionSignal
	}

	// Return an error if any were encountered during execution.
	if errorCounter > 0 {
		return errors.New("errors occurred during execution")
	}

	return nil
}
