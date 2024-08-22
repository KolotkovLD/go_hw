package hw05parallelexecution

import (
	"context"
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m < 0 {
		return ErrErrorsLimitExceeded
	}

	wg := &sync.WaitGroup{}
	errChan := make(chan error, n+m) // Channel for collecting errors

	// Create a context that can be cancelled when the maximum number of errors is reached
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Launch goroutines to execute tasks
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for _, task := range tasks {
				select {
				case <-ctx.Done():
					return
				default:
				}
				if err := task(); err != nil {
					errChan <- err
				}
			}
		}()
	}

	// Wait for all goroutines to finish
	wg.Add(n)
	wg.Wait()

	// Check if we have exceeded the maximum number of errors
	if len(errChan) > m {
		close(errChan)
		return ErrErrorsLimitExceeded
	}

	// Close the channel and return any errors
	close(errChan)
	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}
