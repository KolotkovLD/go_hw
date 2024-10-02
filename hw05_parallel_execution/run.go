package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		m = len(tasks) + 1
	}

	var (
		wg         sync.WaitGroup
		errorCount int
		errorChan  = make(chan error, len(tasks))
		doneChan   = make(chan struct{})
		stopChan   = make(chan struct{})
		taskChan   = make(chan Task, len(tasks))
	)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case task, ok := <-taskChan:
					if !ok {
						return
					}
					if err := task(); err != nil {
						errorChan <- err
					}
				case <-stopChan:
					return
				}
			}
		}()
	}

	// Заполняем канал заданий
	go func() {
		defer close(taskChan)
		for _, task := range tasks {
			select {
			case taskChan <- task:
			case <-stopChan:
				return
			}
		}
	}()

	// Обрабатываем ошибки
	go func() {
		defer close(doneChan)
		for err := range errorChan {
			if err != nil {
				errorCount++
				if errorCount >= m {
					close(stopChan)
					return
				}
			}
		}
	}()

	wg.Wait()
	close(errorChan)
	<-doneChan
	if errorCount >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}
