package hw05parallelexecution

import (
	"errors"
	"log"
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
	stopChan := make(chan struct{})
	taskChan := make(chan Task, len(tasks))

	for i := 0; i < n; i++ {
		wg.Add(1)
		go runTask(&wg, taskChan, errorChan, stopChan, i, &errorCount)
	}
	// Заполняем канал заданий
	go sendTasks(taskChan, tasks, stopChan)

	// Обрабатываем ошибки
	go checkErr(taskChan, errorChan, stopChan, m, &errorCount)

	wg.Wait()
	close(errorChan)
	close(stopChan)

	m32 := int32(m)
	if errorCount >= m32 {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func runTask(wg *sync.WaitGroup, taskChan chan Task, errorChan chan<- error, stopChan chan struct{}, workerID int, errorCount *int32) {
	// Запускает таски из канала taskChan
	defer wg.Done()
	log.Printf("Goroutine %d: started\n", workerID)
	for {
		select {
		case <-stopChan:
			log.Printf("Goroutine %d: stopChan closed, exiting\n", workerID)
			return
		case task, ok := <-taskChan:
			if !ok {
				log.Printf("Goroutine %d: taskChan closed, exiting\n", workerID)
				return
			}
			log.Printf("Goroutine %d: received a task\n", workerID)
			if err := task(); err != nil {
				log.Printf("Goroutine %d: task returned error: %v, errorCount: %d\n", workerID, err, *errorCount)
				errorChan <- err
			}
		}
	}
}

func closeChan(taskChan chan Task) {
	if _, ok := <-taskChan; ok {
		log.Printf("CLOSE CHAN")
		close(taskChan)
	}
}

func sendTasks(taskChan chan Task, tasks []Task, stopChan chan struct{}) {
	// Отправляет таски в канал taskChan
	defer closeChan(taskChan)
	for _, task := range tasks {
		select {
		case taskChan <- task:
		case <-stopChan:
			return
		}
	}
}

func checkErr(taskChan chan Task, errorChan chan error, stopChan chan struct{}, m int, errorCount *int32) {
	// Проверяет количество таков с ошибкой и прерывает работу оставшихся
	for err := range errorChan {
		if err != nil {
			atomic.AddInt32(errorCount, 1)
			m32 := int32(m)
			if *errorCount >= m32 {
				stopChan <- struct{}{}
				closeChan(taskChan)
				return
			}
		}
	}
}
