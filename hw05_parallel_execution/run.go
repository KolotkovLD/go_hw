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
		wg            sync.WaitGroup
		errorCount    atomic.Int32
		runTasksCount atomic.Int32
	)

	if m <= 0 {
		m = len(tasks) + 1
	}

	taskChan := make(chan Task)

	// Заполняем канал заданий
	go sendTasks(taskChan, tasks, &errorCount, &runTasksCount, n, m)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go runTask(&wg, taskChan, &errorCount, &runTasksCount, i, n, m)
	}

	wg.Wait()
	if stopRule(&errorCount, &runTasksCount, n, m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func runTask(wg *sync.WaitGroup,
	taskChan chan Task,
	errorCount *atomic.Int32,
	runTasksCount *atomic.Int32,
	workerID int,
	n int,
	m int,
) {
	// Запускает таски из канала taskChan
	defer wg.Done()
	log.Printf("Goroutine %d: started\n", workerID)
	for {
		task, ok := <-taskChan
		runTasksCount.Add(1)
		log.Printf("Goroutine %d: received a task, runTasksCount: %d\n", workerID, runTasksCount.Load())
		if !ok || stopRule(errorCount, runTasksCount, n, m) {
			log.Printf("Goroutine %d: taskChan closed, exiting\n", workerID)
			return
		}

		if err := task(); err != nil {
			errorCount.Add(1)
			log.Printf("Goroutine %d: task returned error: %v, errorCount: %d\n", workerID, err, errorCount.Load())
		}
	}
}

func sendTasks(taskChan chan Task,
	tasks []Task,
	errorCount *atomic.Int32,
	runTasksCount *atomic.Int32,
	n int,
	m int,
) {
	// Отправляет таски в канал taskChan
	defer func() {
		close(taskChan)
		log.Printf("sendTasks is done!!!!!!!")
	}()
	for _, task := range tasks {
		if stopRule(errorCount, runTasksCount, n, m) {
			log.Printf("     [sendTasks]  errorCount: %d ; runTasksCount: %d\n", errorCount.Load(), runTasksCount.Load())
			return
		}
		taskChan <- task
	}
}

func stopRule(errorCount *atomic.Int32, runTasksCount *atomic.Int32, n, m int) bool {
	if errorCount.Load() >= int32(m) {
		log.Printf(" [stop_rule]  errorCount: %d ; runTasksCount: %d\n", errorCount.Load(), runTasksCount.Load())
		return true
	}
	return false
}
