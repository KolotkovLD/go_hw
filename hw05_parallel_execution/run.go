package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {

	var (
		wg         sync.WaitGroup
		errorCount int32
	)

	if m <= 0 {
		m = len(tasks) + 1
	}

	errorChan := make(chan error, len(tasks))
	doneChan := make(chan struct{})
	stopChan := make(chan struct{})
	taskChan := make(chan Task, len(tasks))

	for i := 0; i < n; i++ {
		wg.Add(1)
		go runTask(&wg, taskChan, errorChan, stopChan)
	}
	// Заполняем канал заданий
	go sendTasks(taskChan, tasks, stopChan)

	// Обрабатываем ошибки
	go checkErr(doneChan, errorChan, stopChan, m, &errorCount)

	wg.Wait()
	close(errorChan)
	<-doneChan
	m32 := int32(m)
	if errorCount > m32 {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func runTask(wg *sync.WaitGroup, taskChan chan Task, errorChan chan<- error, stopChan chan struct{}) {
	// Запускает таски из канала taskChan
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
}

func sendTasks(taskChan chan<- Task, tasks []Task, stopChan chan struct{}) {
	// Отправляет таски в канал taskChan
	defer close(taskChan)
	for _, task := range tasks {
		select {
		case taskChan <- task:
		case <-stopChan:
			return
		}
	}
}

func checkErr(doneChan chan struct{}, errorChan chan error, stopChan chan struct{}, m int, errorCount *int32) {
	// Проверяет количество таков с ошибкой и прерывает работу оставшихся
	defer close(doneChan)
	for err := range errorChan {
		if err != nil {
			atomic.AddInt32(errorCount, 1)
			m32 := int32(m)
			if *errorCount > m32 {
				close(stopChan)
				return
			}
		}
	}
}
